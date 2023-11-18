package entity

type ProviderResponse struct {
	Supplier_inn string `json:"supplier_inn" db:"supplier_inn" binding:"-"`
	Upd          int    `json:"upd" db:"upd" binding:"-"`
	Contract     int    `json:"contract" db:"contract" binding:"-"`
	Facap        int    `json:"facap" db:"facap" binding:"-"`
	DoneContr    int    `json:"donecontr" db:"donecontr" binding:"-"`
	Min          string `json:"min" db:"min" binding:"-"`
	Max          string `json:"max" db:"max" binding:"-"`
	BlockSum     string `json:"blocksum" db:"blocksum" binding:"-"`
}
