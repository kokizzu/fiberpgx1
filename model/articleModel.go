package model

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type ArticleModel struct {
	Id        int64
	Title     string
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeleteAt  sql.NullTime
	CreatedBy int32
}

func (u *ArticleModel) FindById(adapter *Adapter) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	const selectById = `
SELECT title, body, created_at, updated_at, delete_at, created_by
FROM articles
WHERE id = $1
LIMIT 1
`
	err := adapter.QueryRow(ctx, selectById, u.Id).Scan(
		&u.Title,
		&u.Body,
		&u.CreatedAt,
		&u.UpdatedAt,
		&u.DeleteAt,
		&u.CreatedBy,
	)

	if err != nil {
		return fmt.Errorf(`ArticleModel) FindById: %w`, err)
	}
	return nil
}

func (u *ArticleModel) Insert(adapter *Adapter) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	const insert = `
INSERT INTO articles (title, body, created_by)
VALUES ($1, $2, $3)
RETURNING id, created_at, updated_at
`
	err := adapter.QueryRow(ctx, insert, u.Title, u.Body, u.CreatedBy).Scan(
		&u.Id,
		&u.CreatedAt,
		&u.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf(`ArticleModel) Insert: %w`, err)
	}
	return nil
}

func (u *ArticleModel) UpdateById(adapter *Adapter) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	const updateById = `
UPDATE articles
SET title = $1, body = $2, updated_at = $3
WHERE id = $4
RETURNING updated_at
`
	err := adapter.QueryRow(ctx, updateById, u.Title, u.Body, time.Now(), u.Id).Scan(
		&u.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf(`ArticleModel) UpdateById: %w`, err)
	}
	return err
}

func (u *ArticleModel) DeleteById(adapter *Adapter) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	const deleteById = `
UPDATE articles
SET delete_at = $1
WHERE id = $2
RETURNING delete_at
`
	err := adapter.QueryRow(ctx, deleteById, time.Now(), u.Id).Scan(
		&u.DeleteAt,
	)

	if err != nil {
		return fmt.Errorf(`ArticleModel) DeleteById: %w`, err)
	}
	return nil
}

func (u *ArticleModel) RestoreById(adapter *Adapter) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	const restoreById = `
UPDATE articles
SET delete_at = NULL
WHERE id = $1
RETURNING delete_at
`
	err := adapter.QueryRow(ctx, restoreById, u.Id).Scan(
		&u.DeleteAt,
	)

	if err != nil {
		return fmt.Errorf(`ArticleModel) RestoreById: %w`, err)
	}
	return nil
}

func (u *ArticleModel) FindOffsetLimit(adapter *Adapter, offset int, limit int) (res []ArticleModel, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	const selectOffsetLimit = `
SELECT id, title, body, created_at, updated_at, delete_at, created_by
FROM articles
WHERE delete_at IS NULL
ORDER BY id DESC
OFFSET $1 LIMIT $2
`
	rows, err := adapter.Query(ctx, selectOffsetLimit, offset, limit)
	if err != nil {
		return nil, fmt.Errorf(`ArticleModel) FindOffsetLimit: %w`, err)
	}
	defer rows.Close()
	for rows.Next() {
		var row ArticleModel
		err = rows.Scan(
			&row.Id,
			&row.Title,
			&row.Body,
			&row.CreatedAt,
			&row.UpdatedAt,
			&row.DeleteAt,
			&row.CreatedBy,
		)
		if err != nil {
			return nil, fmt.Errorf(`ArticleModel) FindOffsetLimit: %w`, err)
		}
		res = append(res, row)
	}

	return res, nil
}
