package repository

import (
	"context"

	"ecom/internal/domain"
)

type GoodRepo interface {
	GetAllGoods(ctx context.Context, filters []domain.GormFilter, ordersStr string) ([]domain.Good, error)
	GetGoodByID(ctx context.Context, id string) (domain.Good, error)
	AddGood(ctx context.Context, good domain.Good) (string, error)
	UpdateGood(ctx context.Context, good domain.Good) error
	DeleteGood(ctx context.Context, id string) error
}

type AuthRepo interface {
	CheckUserExists(username string) bool
	InsertUser(user domain.UserInfo) error
	GetUserByUserName(username string) (domain.UserInfo, error)
}

type PaginationRepo interface {
	GetRecordsCount(table string) (int, error)
}
