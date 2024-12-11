package inmemory

import (
	"ecom/internal/domain"
	"ecom/internal/errs"
	"ecom/internal/repository"
)

type GoodRepoInMemory struct {
	goods map[string]domain.Good
}

func NewGoodRepoInMemory() repository.GoodRepo {
	return &GoodRepoInMemory{
		goods: make(map[string]domain.Good),
	}
}

func (r *GoodRepoInMemory) GetAllGoods() ([]domain.Good, error) {
	goods := make([]domain.Good, 0, len(r.goods))

	for _, good := range r.goods {
		goods = append(goods, good)
	}

	return goods, nil
}

func (r *GoodRepoInMemory) GetGoodByID(id string) (domain.Good, error) {
	good, exists := r.goods[id]
	if !exists {
		return domain.Good{}, errs.ErrGoodNotFound
	}
	return good, nil
}

func (r *GoodRepoInMemory) AddGood(good domain.Good) (string, error) {
	r.goods[good.ID] = good
	return good.ID, nil
}

func (r *GoodRepoInMemory) UpdateGood(good domain.Good) error {
	_, exists := r.goods[good.ID]
	if !exists {
		return errs.ErrGoodNotFound
	}
	r.goods[good.ID] = good
	return nil
}

func (r *GoodRepoInMemory) DeleteGood(id string) error {
	_, exists := r.goods[id]
	if !exists {
		return errs.ErrGoodNotFound
	}

	delete(r.goods, id)
	return nil
}
