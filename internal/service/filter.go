package service

import (
	"tender/internal/entity"
	"tender/internal/repository"
)

type FilterService struct {
	repo repository.Filter
}

func NewFilterService(repo repository.Filter) *FilterService {
	return &FilterService{
		repo: repo,
	}
}

func (s *FilterService) GetAllKpgz(kpgz string) ([]entity.ProviderResponse, error) {
	return s.repo.GetAllKpgz(kpgz)
}

func (s *FilterService) GetProviderByInn(inn string) (entity.ProviderResponse, error) {
	return s.repo.GetProviderByInn(inn)
}
