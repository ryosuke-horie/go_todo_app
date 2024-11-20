package main

import (
	"fmt"
	"net/http"
	"os"
)

// 以下の問題点を持っているコード
// 1. テスト完了後に終了する方法がない
// 2. 出力が検証しにくい
// 3. 異常時にos.Exitが呼ばれる
// 4. ポート番号が固定されるためテストの起動に失敗する可能性がある

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
