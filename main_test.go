package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"testing"

	"golang.org/x/sync/errgroup"
)

func TestRun(t *testing.T) {
	// localhost:0を指定するとnet/httpパッケージが自動的にポートを割り当てる
	l, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatalf("failed to listen port %v", err)
	}

	// キャンセル可能なcontext.Contextを作成
	ctx, cancel := context.WithCancel(context.Background())

	eg, ctx := errgroup.WithContext(ctx)
	// 別goroutineでテスト対象のrun関数を実行しHTTPサーバーを起動
	eg.Go(func() error {
		return run(ctx, l)
	})

	in := "message"
	// ポートを自動割り当てしたアドレスを使用してURLを作成
	url := fmt.Sprintf("http://%s/%s", l.Addr().String(), in)
	// どんなポート番号でアクセスしているかログで出力
	t.Logf("try request to %s", url)
	// GETリクエストを送信
	rsp, err := http.Get(url)
	if err != nil {
		t.Errorf("failed to get: %+v", err)
	}
	defer rsp.Body.Close()

	got, err := io.ReadAll(rsp.Body)
	if err != nil {
		t.Fatalf("failed to read body: %v", err)
	}

	// httpサーバーの戻り値を検証する
	// 期待する文字列であることを確認する
	want := fmt.Sprintf("Hello, %s", in)
	if string(got) != want {
		t.Errorf("want %q, but got %q", want, got)
	}

	// run関数に終了通知を送る
	cancel()

	// run関数の戻り値を検証する
	if err := eg.Wait(); err != nil {
		t.Fatal(err)
	}
}
