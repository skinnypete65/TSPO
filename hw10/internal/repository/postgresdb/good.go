package postgresdb

import (
	"ecom/internal/domain"
	"ecom/internal/repository"
	"gorm.io/gorm"
)

type GoodPostgresRepo struct {
	db *gorm.DB
}

func NewGoodPostgresRepo(
	db *gorm.DB,
) repository.GoodRepo {
	return &GoodPostgresRepo{
		db: db,
	}
}

func (r *GoodPostgresRepo) GetAllGoods() ([]domain.Good, error) {
	var goods []domain.Good
	err := r.db.Find(&goods).Error
	if err != nil {
		return nil, err
	}
	return goods, nil
}

func (r *GoodPostgresRepo) GetGoodByID(id string) (domain.Good, error) {
	var good domain.Good
	res := r.db.First(&good, "good_id = ?", id)
	if res.Error != nil {
		return domain.Good{}, res.Error
	}
	return good, nil
}
func (r *GoodPostgresRepo) AddGood(good domain.Good) (string, error) {
	err := r.db.Create(&good).Error
	if err != nil {
		return "", err
	}
	return good.ID, nil
}
func (r *GoodPostgresRepo) UpdateGood(good domain.Good) error {
	res := r.db.Model(&domain.Good{}).Where("good_id = ?", good.ID).Updates(good)
	return res.Error
}
func (r *GoodPostgresRepo) DeleteGood(id string) error {
	res := r.db.Delete(&domain.Good{}, "good_id = ?", id)
	return res.Error
}
