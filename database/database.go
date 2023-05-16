package database

import (
	"database/sql"
	"fmt"
	"github.com/Sortren/event-log/pkg/config"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bundebug"

	_ "github.com/uptrace/bun/driver/pgdriver"
)

func Connect(cfg config.Postgres) (*bun.DB, error) {
	sqldb, err := sql.Open("pg", cfg.Dsn())
	if err != nil {
		return nil, fmt.Errorf("can't open database connection, err: %w", err)
	}

	db := bun.NewDB(sqldb, pgdialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	return db, nil
}
