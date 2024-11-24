package handler

import (
	"net/http"

	"github.com/ryosuke-horie/go_todo_app/entity"
	"github.com/ryosuke-horie/go_todo_app/store"
)

// データを格納するstore.TaskStoreをフィールドに持つ構造体
type ListTask struct {
	Store *store.TaskStore
}

type task struct {
	ID     entity.TaskID     `json:"id"`
	Title  string            `json:"title"`
	Status entity.TaskStatus `json:"status"`
}

// タスクのリストを返すハンドラ
func (lt *ListTask) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// リクエストのコンテキストを取得
	ctx := r.Context()

	// 全てのタスクを取得
	tasks := lt.Store.All()

	// JSONフォーマットでレスポンスを返すために、
	// 取得したタスクをtask構造体のスライスに変換
	rsp := []task{}
	for _, t := range tasks {
		rsp = append(rsp, task{
			ID:     t.ID,
			Title:  t.Title,
			Status: t.Status,
		})
	}
	RespondJSON(ctx, w, rsp, http.StatusOK)
}
