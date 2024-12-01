package testutil

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/go-redis/redis/v8"
)

// テスト環境ごとに接続を変更するRedis用テストヘルパー
// テストの実行環境によって異なるRedisの接続情報の差分を吸収する
func OpenRedisForTest(t *testing.T) *redis.Client {
	t.Helper()

	host := "127.0.0.1"
	port := 36379

	// GitHub Actionsでの実行と判断したらportを変更する
	if _, defined := os.LookupEnv("CI"); defined {
		// https://docs.github.com/ja/actions/using-containerized-services/creating-redis-service-containers#configuring-the-runner-job
		port = 6379
	}
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Password: "",
		DB:       0, // default database number
	})
	if err := client.Ping(context.Background()).Err(); err != nil {
		t.Fatalf("failed to connect redis: %s", err)
	}
	return client
}
