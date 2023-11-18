package service

import (
	"tender/internal/entity"
	"tender/internal/repository"
)

type Filter interface {
	GetAllKpgz(kpgz string) ([]entity.ProviderResponse, error)
	GetProviderByInn(inn string) (entity.ProviderResponse, error)
}

type Service struct {
	Filter
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Filter: NewFilterService(repo.Filter),
	}
}
