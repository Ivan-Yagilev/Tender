package repository

import (
	"github.com/jmoiron/sqlx"
	"tender/internal/entity"
)

type Filter interface {
	GetAllKpgz(kpgz string) ([]entity.ProviderResponse, error)
}

type Repository struct {
	Filter
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Filter: NewFilterPostgres(db),
	}
}
