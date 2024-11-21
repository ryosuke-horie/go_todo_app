package main

import "net/http"

// HTTPハンドラーの定義
func NewMux() http.Handler {
	mux := http.NewServeMux()

	// HTTPサーバーが稼働中であることを確認するためのヘルスチェックエンドポイント
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		// 静的解析を通すために明示的に戻り値を捨てる
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	})

	return mux
}
