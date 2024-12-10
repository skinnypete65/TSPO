package rest

import (
	"encoding/json"
	"errors"
	"net/http"

	"ecom/internal/converter"
	"ecom/internal/errs"
	"ecom/internal/response"
	"ecom/internal/service"
	"ecom/internal/transport/rest/dto"
	"github.com/go-playground/validator/v10"
)

const (
	goodIDKey = "good_id"
)

type GoodHandler struct {
	goodService   service.GoodService
	goodConverter converter.GoodConverter
	validate      *validator.Validate
}

func NewGoodHandler(
	goodService service.GoodService,
	goodConverter converter.GoodConverter,
	validate *validator.Validate,
) *GoodHandler {
	return &GoodHandler{
		goodService:   goodService,
		goodConverter: goodConverter,
		validate:      validate,
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
	goods, err := h.goodService.GetAllGoods()

	if err != nil {
		response.InternalServerError(w)
		return
	}

	goodsBytes, err := json.Marshal(h.goodConverter.MapDomainsToDtos(goods))
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

	good, err := h.goodService.GetGoodByID(goodID)
	if err != nil {
		if errors.Is(err, errs.ErrGoodNotFound) {
			response.BadRequest(w, err.Error())
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

	id, err := h.goodService.AddGood(h.goodConverter.MapRequestToDomain(good))
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

	err = h.goodService.UpdateGood(h.goodConverter.MapRequestToDomain(good))
	if err != nil {
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

	err := h.goodService.DeleteGood(goodID)
	if err != nil {
		if errors.Is(err, errs.ErrGoodNotFound) {
			response.BadRequest(w, err.Error())
			return
		}

		response.InternalServerError(w)
		return
	}

	response.OKMessage(w, "Good has been deleted")
}
