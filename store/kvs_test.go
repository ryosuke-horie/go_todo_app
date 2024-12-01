package store

import (
	"context"
	"testing"

	"github.com/ryosuke-horie/go_todo_app/entity"
	"github.com/ryosuke-horie/go_todo_app/testutil"
)

// kvsへの保存用関数をテストする
func TestKVS_Save(t *testing.T) {
	t.Parallel()

	// テスト用にRedisのクライアントを取得
	cli := testutil.OpenRedisForTest(t)
	// kvsに割り当て
	sut := &KVS{Cli: cli}
	key := "TestKVS_Save"
	uid := entity.UserID(1234)
	ctx := context.Background()
	t.Cleanup(func() {
		cli.Del(ctx, key)
	})
	if err := sut.Save(ctx, key, uid); err != nil {
		t.Errorf("want no error, but got %v", err)
	}
}
