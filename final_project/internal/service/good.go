package service

import (
	"ecom/internal/domain"
	"ecom/internal/repository"
	"github.com/google/uuid"
)

type GoodService interface {
	GetAllGoods(filters []domain.GormFilter) ([]domain.Good, error)
	GetGoodByID(id string) (domain.Good, error)
	AddGood(good domain.Good) (string, error)
	UpdateGood(good domain.Good) error
	DeleteGood(id string) error
}

type goodService struct {
	goodRepo repository.GoodRepo
}

func NewGoodService(goodRepo repository.GoodRepo) GoodService {
	return &goodService{
		goodRepo: goodRepo,
	}
}

func (s *goodService) GetAllGoods(filters []domain.GormFilter) ([]domain.Good, error) {
	return s.goodRepo.GetAllGoods(filters)
}

func (s *goodService) GetGoodByID(id string) (domain.Good, error) {
	return s.goodRepo.GetGoodByID(id)
}

func (s *goodService) AddGood(good domain.Good) (string, error) {
	id := uuid.New().String()
	good.ID = id

	return s.goodRepo.AddGood(good)
}

func (s *goodService) UpdateGood(good domain.Good) error {
	return s.goodRepo.UpdateGood(good)
}

func (s *goodService) DeleteGood(id string) error {
	return s.goodRepo.DeleteGood(id)
}
