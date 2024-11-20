package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/sync/errgroup"
)

// メインプロセスを実装するときに資料されるrun関数のパターン
// context.Context方の値から終了通知を待機するrun関数
// チャネルからの終了通知があると次の順序でrun関数が終了する
// 1. 「<-ctx.Done()」の津g期の行の*http.Server.Shutdownメソッドを呼び出す
// 2. 別goroutineで実行していた*http.Server.ListenAndServeメソッドが終了する
// 3. 別goroutineで実行されていた無名関数が終了する
// 4. run関数の最後で待機していたerrgroup.Group.Waitメソッドが終了する
// 5.
func run(ctx context.Context) error {
	s := http.Server{
		Addr: ":18080",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "hello, %s", r.URL.Path[1:])
		}),
	}

	eg, ctx := errgroup.WithContext(ctx)
	// 別のgoroutineでサーバを起動する
	eg.Go(func() error {
		// http.ErrServerClosedは
		// http.Server.shutdownが正常に終了したことを示すため以上ではない
		if err := s.ListenAndServe(); err != nil &&
			err != http.ErrServerClosed {
			log.Printf("failed to close: %+v", err)
			return err
		}

		return nil
	})

	// チャンネルからの通知（終了通知）を待機する
	<-ctx.Done()
	if err := s.Shutdown(context.Background()); err != nil {
		log.Printf("failed to shutdown: %+v", err)
	}
	// Goメソッドで起動した別のgoroutineの終了を待つ
	return eg.Wait()
}

func main() {
	if err := run(context.Background()); err != nil {
		log.Printf("failed to terminate server: %v", err)
	}
}
