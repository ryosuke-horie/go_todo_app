package service

import (
	"context"

	"github.com/ryosuke-horie/go_todo_app/entity"
	"github.com/ryosuke-horie/go_todo_app/store"
)

// storeパッケージへの直接参照を避けるためにインタフェースを定義する（DI）

//go:generate go run github.com/matryer/moq -out maq_test.go . TaskAdder TaskLister
type TaskAdder interface {
	AddTask(ctx context.Context, db store.Execer, t *entity.Task) error
}

type TaskLister interface {
	ListTasks(ctx context.Context, db store.Queryer) (entity.Tasks, error)
}
