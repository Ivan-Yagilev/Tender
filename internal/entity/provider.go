package entity

import "time"

type ProviderResponse struct {
	Inn            string  `json:"inn"`
	FailedDedlines float32 `json:"failed_dedlines"`
	AvgUdpContract float32 `json:"avg_udp_contract"`
	Activity       float32 `json:"activity"`
	Total          float32 `json:"total"`
}

type SingleProviderResponse struct {
	Inn            string  `json:"inn"`
	FailedDedlines float32 `json:"failed_dedlines"`
	AvgUdpContract float32 `json:"avg_udp_contract"`
	Activity       float32 `json:"activity"`
	Total          float32 `json:"total"`
}

type ProviderDb struct {
	Supplier_inn string    `db:"supplier_inn"`
	Upd          int       `db:"upd"`
	Contract     int       `db:"contract"`
	Facap        int       `db:"facap"`
	DoneContr    int       `db:"donecontr"`
	Min          time.Time `db:"min"`
	Max          time.Time `db:"max"`
	BlockSum     string    `db:"blocksum"`
	//BlockInn     string    `db:"blinn"`
	//ParticipantInn string `json:"participant_inn" db:"participant_inn"`
	//Kpgz           string `json:"kpgz" db:"kpgz""`
}
