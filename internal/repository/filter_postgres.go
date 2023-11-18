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

	query := fmt.Sprintf("SELECT price, kpgz FROM %s WHERE $1 LIKE kpgz;", ksTable)
	err := r.db.Select(&lists, query, kpgz)

	if err != nil {
		return nil, nil
	}

	return lists, err
}

func (r *FilterPostgres) GetProviderByInn(inn string) (entity.ProviderResponse, error) {
	var provider entity.ProviderResponse

	query := fmt.Sprintf("%s", ksTable)
	err := r.db.Select(&provider, query, inn)

	if err != nil {
		return provider, nil
	}

	return provider, err
}
