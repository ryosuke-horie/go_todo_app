package fixture

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/ryosuke-horie/go_todo_app/entity"
)

func User(u *entity.User) *entity.User {
	// 新しいentity.Userオブジェクトを生成する
	result := &entity.User{
		ID:       entity.UserID(rand.Int()),
		Name:     "test" + strconv.Itoa(rand.Int())[:5],
		Password: "password",
		Role:     "admin",
		Created:  time.Now(),
		Modified: time.Now(),
	}
	// ユーザー情報を生成しない場合にはそのまま返す
	if u == nil {
		return result
	}

	// カスタムの値がある場合にはそれを設定する
	if u.ID != 0 {
		result.ID = u.ID
	}
	if u.Name != "" {
		result.Name = u.Name
	}
	if u.Password != "" {
		result.Password = u.Password
	}
	if u.Role != "" {
		result.Role = u.Role
	}
	if !u.Created.IsZero() {
		result.Created = u.Created
	}
	if !u.Modified.IsZero() {
		result.Modified = u.Modified
	}
	return result
}
