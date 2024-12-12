package postgresdb

import (
	"fmt"

	"ecom/internal/repository"
	"gorm.io/gorm"
)

type paginationRepoPostgres struct {
	db *gorm.DB
}

func NewPaginationRepoPostgres(db *gorm.DB) repository.PaginationRepo {
	return &paginationRepoPostgres{
		db: db,
	}
}

func (r *paginationRepoPostgres) GetRecordsCount(table string) (int, error) {
	count := int64(0)

	sqlTableQuery := fmt.Sprintf("SELECT count(*) FROM %s", table)
	err := r.db.Raw(sqlTableQuery).Count(&count).Error

	return int(count), err
}
