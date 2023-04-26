package model

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type Adapter struct {
	*pgxpool.Pool
}
