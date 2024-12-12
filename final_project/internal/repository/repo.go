package repository

import "ecom/internal/domain"

type GoodRepo interface {
	GetAllGoods(filters []domain.GormFilter) ([]domain.Good, error)
	GetGoodByID(id string) (domain.Good, error)
	AddGood(good domain.Good) (string, error)
	UpdateGood(good domain.Good) error
	DeleteGood(id string) error
}

type AuthRepo interface {
	CheckUserExists(username string) bool
	InsertUser(user domain.UserInfo) error
	GetUserByUserName(username string) (domain.UserInfo, error)
}

type PaginationRepo interface {
	GetRecordsCount(table string) (int, error)
}
