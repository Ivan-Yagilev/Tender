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
	listDb, err := s.repo.GetAllKpgz(kpgz)
	if err != nil {
		return nil, nil
	}

	lists := make([]entity.ProviderResponse, len(listDb))

	for i := 0; i < len(listDb); i++ {
		lists[i].Inn = listDb[i].Supplier_inn
		lists[i].FailedDedlines = float32(listDb[i].Facap) / float32(listDb[i].Upd)
		lists[i].AvgUdpContract = float32(listDb[i].Upd) / float32(listDb[i].Contract)

		DayDiff := (listDb[i].Max.Sub(listDb[i].Min) / 24).Hours()

		if DayDiff == 0 {
			lists[i].Activity = float32(listDb[i].DoneContr)
		} else {
			lists[i].Activity = float32(listDb[i].DoneContr) / float32(DayDiff)
		}

		lists[i].Total = 0.674*lists[i].FailedDedlines + 0.523*lists[i].AvgUdpContract + 0.343*lists[i].Activity
	}

	return lists, nil
}

func (s *FilterService) GetProviderByInn(inn string) (entity.ProviderResponse, error) {
	return s.repo.GetProviderByInn(inn)
}
