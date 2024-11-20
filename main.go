package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/sync/errgroup"
)

// メインプロセスを実装するときに資料されるrun関数のパターン
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
		fmt.Fprintf("failed to terminate server: %v", err)
	}
}
