package database

import (
	"github.com/rs/zerolog"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	Pool *pgxpool.Pool
	log  *zerolog.Logger
}