package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/ryosuke-horie/go_todo_app/handler"
	"github.com/ryosuke-horie/go_todo_app/store"
)

// HTTPハンドラーの定義
func NewMux() http.Handler {
	// mux := http.NewServeMux()
	mux := chi.NewRouter()

	// HTTPサーバーが稼働中であることを確認するためのヘルスチェックエンドポイント
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		// 静的解析を通すために明示的に戻り値を捨てる
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	})

	v := validator.New()
	at := &handler.AddTask{Store: store.Tasks, Validator: v}
	mux.Post("/tasks", at.ServeHTTP)

	lt := &handler.ListTask{Store: store.Tasks}
	mux.Get("/tasks", lt.ServeHTTP)

	return mux
}
