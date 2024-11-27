package service

import (
	"context"
	"fmt"

	"github.com/ryosuke-horie/go_todo_app/entity"
	"github.com/ryosuke-horie/go_todo_app/store"
)

type ListTasks struct {
	DB   store.Queryer
	Repo TaskLister
}

func (l *ListTasks) ListTasks(ctx context.Context) (entity.Tasks, error) {
	ts, err := l.Repo.ListTasks(ctx, l.DB)
	if err != nil {
		return nil, fmt.Errorf("failed to list:%w", err)
	}

	return ts, nil
}
