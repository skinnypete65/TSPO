package service

import (
	"context"

	"ecom/internal/domain"
	"ecom/internal/repository"
	"github.com/google/uuid"
)

type GoodService interface {
	GetAllGoods(ctx context.Context, filters []domain.GormFilter, ordersStr string) ([]domain.Good, error)
	GetGoodByID(ctx context.Context, id string) (domain.Good, error)
	AddGood(ctx context.Context, good domain.Good) (string, error)
	UpdateGood(ctx context.Context, good domain.Good) error
	DeleteGood(ctx context.Context, id string) error
}

type goodService struct {
	goodRepo repository.GoodRepo
}

func NewGoodService(goodRepo repository.GoodRepo) GoodService {
	return &goodService{
		goodRepo: goodRepo,
	}
}

func (s *goodService) GetAllGoods(ctx context.Context, filters []domain.GormFilter, ordersStr string) ([]domain.Good, error) {
	return s.goodRepo.GetAllGoods(ctx, filters, ordersStr)
}

func (s *goodService) GetGoodByID(ctx context.Context, id string) (domain.Good, error) {
	return s.goodRepo.GetGoodByID(ctx, id)
}

func (s *goodService) AddGood(ctx context.Context, good domain.Good) (string, error) {
	id := uuid.New().String()
	good.ID = id

	return s.goodRepo.AddGood(ctx, good)
}

func (s *goodService) UpdateGood(ctx context.Context, good domain.Good) error {
	return s.goodRepo.UpdateGood(ctx, good)
}

func (s *goodService) DeleteGood(ctx context.Context, id string) error {
	return s.goodRepo.DeleteGood(ctx, id)
}
