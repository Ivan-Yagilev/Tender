package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"tender/internal/entity"
)

type FilterPostgres struct {
	db *sqlx.DB
}

func NewFilterPostgres(db *sqlx.DB) *FilterPostgres {
	return &FilterPostgres{
		db: db,
	}
}

func (r *FilterPostgres) GetAllKpgz(kpgz string) ([]entity.ProviderResponse, error) {
	var lists []entity.ProviderResponse

	query := fmt.Sprintf("SELECT kpgz FROM %s WHERE $1 LIKE kpgz;", ksTable)
	err := r.db.Select(&lists, query, kpgz)

	if err != nil {
		return nil, nil
	}

	return lists, err
}
