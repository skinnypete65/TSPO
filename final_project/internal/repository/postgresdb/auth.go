package postgresdb

import (
	"ecom/internal/domain"
	"ecom/internal/repository"
	"gorm.io/gorm"
)

type AuthRepoPostgres struct {
	db *gorm.DB
}

func NewAuthRepoPostgres(db *gorm.DB) repository.AuthRepo {
	return &AuthRepoPostgres{
		db: db,
	}
}

func (r *AuthRepoPostgres) CheckUserExists(username string) bool {
	cnt := int64(0)

	r.db.Model(&domain.UserInfo{}).
		Where("username = ?", username).
		Count(&cnt)

	return cnt > 0
}

func (r *AuthRepoPostgres) InsertUser(user domain.UserInfo) error {
	return r.db.Create(&user).Error
}

func (r *AuthRepoPostgres) GetUserByUserName(username string) (domain.UserInfo, error) {
	user := domain.UserInfo{}
	err := r.db.Where("username = ?", username).First(&user).Error

	return user, err
}
