package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// ListenAndServeは第一引数のアドレス文字列と第二引数のハンドラでHTTPサーバを起動する
	err := http.ListenAndServe(
		":18080",
		// 第二引数に単一のハンドラを渡すことで、全てのリクエストに対して同じ処理を行う
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// パスを使ってレスポンスを返す
			fmt.Fprintf(w, "hello, %s", r.URL.Path[1:])
		}),
	)

	if err != nil {
		fmt.Printf("failed to terminate server: %v", err)
		os.Exit(1)
	}

}
