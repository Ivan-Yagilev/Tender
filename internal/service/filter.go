package service

import (
	"fmt"
	"math"
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
			lists[i].Activity = float32(math.Round((float64(listDb[i].DoneContr) * 100) / 100))
		} else {
			lists[i].Activity = float32(math.Round((float64(listDb[i].DoneContr)/float64(DayDiff))*100) / 100)
		}

		var total1, total2, total3 float64

		if lists[i].FailedDedlines == 0 {
			total1 = 0
		} else {
			total1 = -math.Log(float64(lists[i].FailedDedlines))
		}

		if lists[i].Activity == 0 {
			total2 = 0
		} else {
			total2 = -math.Log(float64(lists[i].Activity))
		}

		if lists[i].AvgUdpContract == 0 {
			total3 = 0
		} else {
			total3 = math.Log(float64(1 / lists[i].AvgUdpContract))
		}

		lists[i].Total = float32(math.Round((total1+total2+total3)*100) / 100)
	}

	return lists, nil
}

func (s *FilterService) GetProviderByInn(inn int) (entity.SingleProviderResponse, error) {
	var (
		Provider entity.SingleProviderResponse
	)
	tmp, err := s.repo.GetProviderByInn(inn)
	fmt.Println(tmp, err)

	return Provider, err
}
