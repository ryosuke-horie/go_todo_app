package auth

import (
	"context"
	_ "embed"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/ryosuke-horie/go_todo_app/clock"
	"github.com/ryosuke-horie/go_todo_app/entity"
)

// ↓バイト配列で鍵を読み込む

//go:embed cert/secret.pem
var rawPrivKey []byte

//go:embed cert/public.pem
var rawPubKey []byte

// ファイルを読み込んで鍵としてデータを保持する型
type JWTer struct {
	PrivateKey, PublicKey jwk.Key
	Store                 Store
	Clocker               clock.Clocker
}

// 生成したJWTをKVストアに保存するためのインターフェース
//
//go:generate go run github.com/matryer/moq -out moq_test.go . Store
type Store interface {
	Save(ctx context.Context, key string, userID entity.UserID) error
	Load(ctx context.Context, key string) (entity.UserID, error)
}

func NewJWTer(s Store, c clock.Clocker) (*JWTer, error) {
	j := &JWTer{Store: s}
	privkey, err := parse(rawPrivKey)
	if err != nil {
		return nil, fmt.Errorf("failed in NewJWTer: private key: %w", err)
	}
	pubkey, err := parse(rawPubKey)
	if err != nil {
		return nil, fmt.Errorf("failed in NewJWTer: public key: %w", err)
	}
	j.PrivateKey = privkey
	j.PublicKey = pubkey
	j.Clocker = c
	return j, nil
}

// バイト配列から鍵をパースする
func parse(rawKey []byte) (jwk.Key, error) {
	// パッケージを利用して鍵情報を含むバイト配列から
	// jwxパッケージで利用可能なjwk.Keyインターフェースを満たす型の値を生成する
	key, err := jwk.ParseKey(rawKey, jwk.WithPEM(true))
	if err != nil {
		return nil, err
	}
	return key, nil
}

const (
	RoleKey     = "role"
	UserNameKey = "user_name"
)

// 署名済みのJWTを生成する
// 引数に渡したユーザーの値に対してJWTを発行する
func (j *JWTer) GenerateToken(ctx context.Context, u entity.User) ([]byte, error) {
	tok, err := jwt.NewBuilder().
		JwtID(uuid.New().String()).
		Issuer(`github.com/ryosuke-horie/go_todo_app`).
		Subject("access_token").
		IssuedAt(j.Clocker.Now()).
		// redisのexpireはこれを使う。
		// https://pkg.go.dev/github.com/go-redis/redis/v8#Client.Set
		// clock.Durationだから Subする必要がある
		Expiration(j.Clocker.Now().Add(30*time.Minute)).
		Claim(RoleKey, u.Role).
		Claim(UserNameKey, u.Name).
		Build()
	if err != nil {
		return nil, fmt.Errorf("GenerateToken: failed to build token: %w", err)
	}
	// jwtをRedisに保存して管理
	if err := j.Store.Save(ctx, tok.JwtID(), u.ID); err != nil {
		return nil, err
	}

	// 秘密鍵でJWTに署名してトークン文字列を生成
	signed, err := jwt.Sign(tok, jwt.WithKey(jwa.RS256, j.PrivateKey))
	if err != nil {
		return nil, err
	}
	return signed, nil
}

// HTTPリクエストのAuthorizationリクエストヘッダにJWTが付与されていることを想定
func (j *JWTer) GetToken(ctx context.Context, r *http.Request) (jwt.Token, error) {
	// HTTPリクエストからJWTである`jwt.Token`インターフェースを満たす型の値を取得
	token, err := jwt.ParseRequest(
		r,
		jwt.WithKey(jwa.RS256, j.PublicKey), // 鍵の指定
		jwt.WithValidate(false),             // 検証は無視（DIしている*auth.JWTer.Clockerフィールドベースで検証するため）
	)
	if err != nil {
		return nil, err
	}

	if err := jwt.Validate(token, jwt.WithClock(j.Clocker)); err != nil {
		return nil, fmt.Errorf("GetToken: failed to validate token: %w", err)
	}

	// redisから削除して手動でexpireさせていることもありうる
	if _, err := j.Store.Load(ctx, token.JwtID()); err != nil {
		return nil, fmt.Errorf("GetToken: %q expired: %w", token.JwtID(), err)
	}

	return token, nil
}
