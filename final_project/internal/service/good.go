package service

import (
	"context"
	"errors"

	"ecom/internal/domain"
	"ecom/internal/errs"
	"ecom/internal/repository"
	"github.com/google/uuid"
	"gorm.io/gorm"
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
	good, err := s.goodRepo.GetGoodByID(ctx, id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return domain.Good{}, errs.ErrGoodNotFound
	}

	return good, err
}

func (s *goodService) AddGood(ctx context.Context, good domain.Good) (string, error) {
	id := uuid.New().String()
	good.ID = id

	ansID, err := s.goodRepo.AddGood(ctx, good)

	return ansID, err
}

func (s *goodService) UpdateGood(ctx context.Context, good domain.Good) error {
	err := s.goodRepo.UpdateGood(ctx, good)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errs.ErrGoodNotFound
	}
	return err
}

func (s *goodService) DeleteGood(ctx context.Context, id string) error {
	err := s.goodRepo.DeleteGood(ctx, id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errs.ErrGoodNotFound
	}
	return err
}
