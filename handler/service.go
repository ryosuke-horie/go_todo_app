package handler

import (
	"context"

	"github.com/ryosuke-horie/go_todo_app/entity"
)

// リクエストの解釈とレスポンスを組み立てる処理をここに集約する
// 構造体や関数ではなくInterfaceな理由↓
// 他パッケージへの参照を取り除いて疎結合にするため
// インタフェースを介して特定の型に依存させないことでモックに処理を入れ替えたテストを書きやすくするため

//go:generate go run github.com/matryer/moq -out moq_test.go . ListTasksService AddTaskService RegisterUserService LoginService
type ListTaskService interface {
	ListTasks(ctx context.Context) (entity.Tasks, error)
}

type AddTaskService interface {
	AddTask(ctx context.Context, title string) (*entity.Task, error)
}
