package store

import (
	"context"

	"github.com/ryosuke-horie/go_todo_app/entity"
)

// taskテーブルへの操作を実装する

// 全てのタスクを取得する
func (r *Repository) ListTask(
	ctx context.Context, db Queryer,
) (entity.Tasks, error) {
	tasks := entity.Tasks{}
	sql := `SELECT
		id, title, status, created, modified
		FROM task;`

	if err := db.SelectContext(ctx, &tasks, sql); err != nil {
		return nil, err
	}
	return tasks, nil
}

// タスクを保存するメソッド
// `*entity.Task`型の値のIDフィールドを更新することによって呼び出し元に発行されたIDを伝える
func (r *Repository) AddTask(
	ctx context.Context, db Execer, t *entity.Task,
) error {
	t.Created = r.Clocker.Now()
	t.Modified = r.Clocker.Now()
	sql := `INSERT INTO task
		(title, status, created, modified)
		VALUES (?,?,?,?)`

	result, err := db.ExecContext(ctx, sql, t.Title, t.Status, t.Created, t.Modified)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	t.ID = entity.TaskID(id)
	return nil
}
