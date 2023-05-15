package database

import (
	"database/sql"
	"fmt"
	"github.com/Sortren/event-log/pkg/config"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

func Connect(cfg config.Postgres) (*bun.DB, error) {
	sqldb, err := sql.Open("pg", cfg.Dsn())
	if err != nil {
		return nil, fmt.Errorf("can't open database connection, err: %w", err)
	}

	return bun.NewDB(sqldb, pgdialect.New()), nil
}
