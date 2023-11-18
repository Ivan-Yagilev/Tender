package entity

type ProviderResponse struct {
	Price float32 `json:"price" db:"price" binding:"required"`
	Kpgz  string  `json:"kpgz" db:"kpgz" binding:"required"`
}
