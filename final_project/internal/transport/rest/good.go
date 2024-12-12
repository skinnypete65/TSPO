package rest

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"ecom/internal/converter"
	"ecom/internal/domain"
	"ecom/internal/errs"
	"ecom/internal/response"
	"ecom/internal/service"
	"ecom/internal/transport/rest/dto"
	"github.com/go-playground/validator/v10"
)

const (
	goodIDKey      = "good_id"
	defaultPage    = 1
	defaultLimit   = 10
	pageKey        = "page"
	limitKey       = "limit"
	minPriceKey    = "min_price"
	maxPriceKey    = "max_price"
	minStockCntKey = "min_stock_cnt"
	measureUnitKey = "measure_unit"
	sortKey        = "sort"
	goodsTableName = "goods"
)

type GoodHandler struct {
	goodService       service.GoodService
	paginationService service.PaginationService
	goodConverter     converter.GoodConverter
	validate          *validator.Validate
}

func NewGoodHandler(
	goodService service.GoodService,
	paginationService service.PaginationService,
	goodConverter converter.GoodConverter,
	validate *validator.Validate,
) *GoodHandler {
	return &GoodHandler{
		goodService:       goodService,
		paginationService: paginationService,
		goodConverter:     goodConverter,
		validate:          validate,
	}
}

// GetAllGoods docs
//
//	@Summary		Получение списка товаров
//	@Tags			goods
//	@Description	Возвращает список всех товаров
//	@ID				get-all-goods
//	@Produce		json
//	@Success		200	{object}	[]dto.Good
//	@Failure		500	{object}	response.Body
//	@Router			/goods [get]
func (h *GoodHandler) GetAllGoods(w http.ResponseWriter, r *http.Request) {
	page, err := h.parseQueryParamInt(r, pageKey, defaultPage)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	limit, err := h.parseQueryParamInt(r, limitKey, defaultLimit)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	filters := make([]domain.GormFilter, 0)

	minPrice, err := h.parseQueryParamInt(r, minPriceKey, 0)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	if minPrice > 0 {
		filters = append(filters, domain.GormFilter{
			Query:  "price >= ?",
			Params: []any{minPrice},
		})
	}

	maxPrice, err := h.parseQueryParamInt(r, maxPriceKey, 0)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}
	if maxPrice > 0 {
		filters = append(filters, domain.GormFilter{
			Query:  "price <= ?",
			Params: []any{maxPrice},
		})
	}

	minStockCnt, err := h.parseQueryParamInt(r, minStockCntKey, 0)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}
	if minStockCnt > 0 {
		filters = append(filters, domain.GormFilter{
			Query:  "stock_quantity > ?",
			Params: []any{minStockCnt},
		})
	}
	measureUnit := r.URL.Query().Get(measureUnitKey)
	if measureUnit != "" {
		filters = append(filters, domain.GormFilter{
			Query:  "measure_unit = ?",
			Params: []any{measureUnit},
		})
	}

	ordersStr := r.URL.Query().Get(sortKey)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	goods, err := h.goodService.GetAllGoods(ctx, filters, ordersStr)

	if err != nil {
		response.InternalServerError(w)
		return
	}

	paginationParams := domain.PaginationParams{
		Page:  page,
		Limit: limit,
	}

	pagination, err := h.paginationService.GetPaginationInfo(goodsTableName, paginationParams)
	if err != nil {
		response.InternalServerError(w)
		return
	}

	goodsInfo := dto.GoodsInfo{
		Goods:      h.goodConverter.MapDomainsToDtos(goods),
		Pagination: pagination,
	}

	goodsBytes, err := json.Marshal(goodsInfo)
	if err != nil {
		response.InternalServerError(w)
		return
	}

	response.WriteResponse(w, http.StatusOK, goodsBytes)
}

// GetGoodByID docs
//
//	@Summary		Получение товара по его айди
//	@Tags			goods
//	@Description	Возвращает товар по его айди
//	@ID				get-good-by-id
//	@Produce		json
//	@Success		200	{object}	dto.Good
//	@Failure		400	{object}	response.Body
//	@Failure		500	{object}	response.Body
//	@Router			/goods/{good_id} [get]
func (h *GoodHandler) GetGoodByID(w http.ResponseWriter, r *http.Request) {
	goodID := r.PathValue(goodIDKey)
	if goodID == "" {
		response.BadRequest(w, "You do not provide good_id")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	good, err := h.goodService.GetGoodByID(ctx, goodID)
	if err != nil {
		if errors.Is(err, errs.ErrGoodNotFound) {
			response.NotFound(w, err.Error())
			return
		}

		response.InternalServerError(w)
		return
	}

	goodBytes, err := json.Marshal(h.goodConverter.MapDomainToDto(good))
	if err != nil {
		response.InternalServerError(w)
		return
	}

	response.WriteResponse(w, http.StatusOK, goodBytes)
}

// AddGood docs
//
//	@Summary		Добавление товара
//	@Tags			goods
//	@Description	Добавление товара
//	@ID				add-good
//	@Param			input	body	dto.Good	true	"Товар"
//	@Produce		json
//	@Success		200	{object}	response.IDResponse
//	@Failure		400	{object}	response.Body
//	@Failure		500	{object}	response.Body
//	@Router			/goods [post]
func (h *GoodHandler) AddGood(w http.ResponseWriter, r *http.Request) {
	var good dto.Good
	err := json.NewDecoder(r.Body).Decode(&good)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	err = h.validate.Struct(good)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	id, err := h.goodService.AddGood(ctx, h.goodConverter.MapRequestToDomain(good))
	if err != nil {
		response.InternalServerError(w)
		return
	}

	response.IdResponse(w, id)
}

// UpdateGood docs
//
//	@Summary		Обновление данных о товаре
//	@Tags			goods
//	@Description	Обновление данных о товаре
//	@ID				update-good
//	@Param			input	body	dto.Good	true	"Товар"
//	@Produce		json
//	@Success		200	{object}	response.Body
//	@Failure		400	{object}	response.Body
//	@Failure		500	{object}	response.Body
//	@Router			/goods/{good_id} [put]
func (h *GoodHandler) UpdateGood(w http.ResponseWriter, r *http.Request) {
	var good dto.Good
	err := json.NewDecoder(r.Body).Decode(&good)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	err = h.validate.Struct(good)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err = h.goodService.UpdateGood(ctx, h.goodConverter.MapRequestToDomain(good))
	if err != nil {
		if errors.Is(err, errs.ErrGoodNotFound) {
			response.NotFound(w, err.Error())
			return
		}
		response.InternalServerError(w)
		return
	}

	response.OKMessage(w, "Good has been updated")
}

// DeleteGoodByID docs
//
//	@Summary		Удаление товара по его айди
//	@Tags			goods
//	@Description	Удаление товара по его айди
//	@ID				delete-good
//	@Produce		json
//	@Success		200	{object}	response.Body
//	@Failure		400	{object}	response.Body
//	@Failure		500	{object}	response.Body
//	@Router			/goods/{good_id} [delete]
func (h *GoodHandler) DeleteGoodByID(w http.ResponseWriter, r *http.Request) {
	goodID := r.PathValue(goodIDKey)
	if goodID == "" {
		response.BadRequest(w, "You do not provide good_id")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := h.goodService.DeleteGood(ctx, goodID)
	if err != nil {
		if errors.Is(err, errs.ErrGoodNotFound) {
			response.NotFound(w, err.Error())
			return
		}

		response.InternalServerError(w)
		return
	}

	response.OKMessage(w, "Good has been deleted")
}

func (h *GoodHandler) parseQueryParamInt(r *http.Request, key string, defaultValue int) (int, error) {
	queryParam := r.URL.Query().Get(key)

	if queryParam == "" {
		return defaultValue, nil
	}

	param, err := strconv.Atoi(queryParam)
	if err != nil {
		return 0, err
	}

	if param == 0 {
		return defaultValue, nil
	}
	return param, nil

}
