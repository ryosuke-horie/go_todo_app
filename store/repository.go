package store

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/ryosuke-horie/go_todo_app/config"
)

func New(ctx context.Context, cfg *config.Config) (*sqlx.DB, func(), error) {
	// sqlx.Connectを使うと内部でpingする
	db, err := sql.Open("mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?parseTime=true",
			cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.Port, cfg.DBName,
		),
	)
	if err != nil {
		return nil, nil, err
	}

	// Openは実際に接続テストが行われない
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, func() { _ = db.Close() }, nil
	}

	xdb := sqlx.NewDb(db, "mysql")
	return xdb, func() { _ = db.Close() }, nil
}
