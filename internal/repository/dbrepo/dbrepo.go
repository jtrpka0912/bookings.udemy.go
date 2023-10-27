package dbrepo

import (
	"database/sql"

	"github.com/jtrpka0912/bookings.udemy.go/internal/config"
	"github.com/jtrpka0912/bookings.udemy.go/internal/repository"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}
}
