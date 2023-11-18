package entity

type ProviderResponse struct {
	Kpgz string `json:"kpgz" db:"kpgz" binding:"required"`
}
