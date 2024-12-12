package postgresdb

import (
	"context"
	"log"

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

func (r *GoodPostgresRepo) GetAllGoods(ctx context.Context, filters []domain.GormFilter, ordersStr string) ([]domain.Good, error) {
	var goods []domain.Good

	tx := r.db.WithContext(ctx)
	for _, filter := range filters {
		tx = tx.Where(filter.Query, filter.Params)
	}

	if ordersStr != "" {
		tx.Order(ordersStr)
	}
	err := tx.Find(&goods).Error
	if err != nil {
		return nil, err
	}
	return goods, nil
}

func (r *GoodPostgresRepo) GetGoodByID(ctx context.Context, id string) (domain.Good, error) {
	var good domain.Good
	res := r.db.WithContext(ctx).First(&good, "good_id = ?", id)
	if res.Error != nil {
		return domain.Good{}, res.Error
	}
	return good, nil
}
func (r *GoodPostgresRepo) AddGood(ctx context.Context, good domain.Good) (string, error) {
	err := r.db.WithContext(ctx).Create(&good).Error
	if err != nil {
		return "", err
	}
	return good.ID, nil
}
func (r *GoodPostgresRepo) UpdateGood(ctx context.Context, good domain.Good) error {
	tx := r.db.WithContext(ctx).Begin()

	log.Println("Start update transaction")
	err := tx.Model(&domain.Good{}).Where("good_id = ?", good.ID).Updates(good).Error
	if err != nil {
		log.Printf("Error happened: %v\nRollback transaction\n", err)
		tx.Rollback()
	} else {
		log.Println("Transaction finished successfully")
	}

	return err
}
func (r *GoodPostgresRepo) DeleteGood(ctx context.Context, id string) error {
	res := r.db.WithContext(ctx).Delete(&domain.Good{}, "good_id = ?", id)
	return res.Error
}
