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
	var (
		lists []entity.ProviderResponse
	)

	query := fmt.Sprintf("SELECT * FROM ( SELECT DISTINCT supplier_inn, count(DISTINCT upd_id) as upd, count(DISTINCT contract_id) AS contract, count(CASE WHEN scheduled_delivery_date < actual_delivery_date AND scheduled_delivery_date != '1970-01-01 00:00:00.000000' THEN supplier_inn END) AS facap FROM contreactexec group by supplier_inn) as t1 LEFT JOIN (SELECT DISTINCT supplier_inn, count(*) AS doneContr, min(conclusion_date), max(conclusion_date) FROM contracts GROUP BY supplier_inn) as t2 on t1.supplier_inn = t2.supplier_inn LEFT JOIN (SELECT DISTINCT blocking.supplier_inn as blInn, sum(blocking.blocking_end_date-blocking.blocking_start_date) as blockSum FROM blocking JOIN contracts on blocking.supplier_inn = contracts.supplier_inn GROUP BY blocking.supplier_inn) as t3 ON t2.supplier_inn = t3.blInn LEFT JOIN (SELECT participant_inn, kpgz FROM ks) as bruh on t3.blInn = bruh.participant_inn WHERE kpgz LIKE '$1%'")
	err := r.db.Select(&lists, query, kpgz)

	fmt.Println(lists)

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
