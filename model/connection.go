package model

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectPostgres(user string, pass string, host string, port int, dbname string) (*pgxpool.Pool, error) {
	const connTpl = `postgres://%s:%s@%s:%d/%s?sslmode=disable&pool_max_conns=%d`
	connStr := fmt.Sprintf(connTpl,
		user,
		pass,
		host,
		port,
		dbname, // or postgres
		32,
	)

	ctx := context.Background()

	return pgxpool.New(ctx, connStr)
}
