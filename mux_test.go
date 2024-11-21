package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewMux(t *testing.T) {
	// これらを利用するとHTTPサーバを起動せずともHTTPハンドラのテストができる
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/health", nil)

	sut := NewMux()
	sut.ServeHTTP(w, r)

	resp := w.Result()
	t.Cleanup(func() {
		_ = resp.Body.Close()
	})

	// ステータスコードが200であることを確認
	if resp.StatusCode != http.StatusOK {
		t.Error("want status code 200, but ", resp.StatusCode)
	}

	// レスポンスボディの内容を確認
	got, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("failed to read body: ", err)
	}

	// レスポンスが期待する文字列であることを確認
	want := `{"status":"ok"}`
	if string(got) != want {
		t.Errorf("want %q, but got %q", want, got)
	}
}
