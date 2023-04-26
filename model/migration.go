package model

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func DoMigrate(pg *pgxpool.Pool) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	const createArticleTable = `
CREATE TABLE IF NOT EXISTS articles (
	id BIGSERIAL PRIMARY KEY NOT NULL
	, title VARCHAR(255) NOT NULL
	, body TEXT NOT NULL
	, created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	, updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	, deleted_at TIMESTAMP NULL
)`
	_, err := pg.Exec(ctx, createArticleTable)
	if err != nil {
		log.Fatalf(`createArticleTable %v`, err)
	}
	const alterArticleTable = `
ALTER TABLE articles ADD COLUMN IF NOT EXISTS created_by INT 
`
	_, err = pg.Exec(ctx, alterArticleTable)
	if err != nil {
		log.Fatalf(`alterArticleTable %v`, err)
	}
}
