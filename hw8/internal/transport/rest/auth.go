package rest

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"ecom/internal/domain"
	"ecom/internal/errs"
	"ecom/internal/response"
	"ecom/internal/service"
	"ecom/internal/transport/rest/dto"
	"github.com/go-playground/validator/v10"
)

type AuthHandler struct {
	validate    *validator.Validate
	authService service.AuthService
}

func NewAuthHandler(
	authService service.AuthService,
	validate *validator.Validate,
) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		validate:    validate,
	}
}

// SignUp docs
//
//	@Summary		Регистрация пользователя
//	@Tags			auth
//	@Description	Регистрация пользователя по логину и паролю
//	@ID				auth-sign-up
//	@Accept			json
//	@Produce		json
//	@Param			input		body		dto.SignUpInput	true	"Логин, пароль, роль"
//	@Success		200			{object}	response.Body
//	@Failure		400,401,409	{object}	response.Body
//	@Failure		500			{object}	response.Body
//	@Failure		default		{object}	response.Body
//	@Router			/auth/sign_up [post]
func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var input dto.SignUpInput

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	err = h.validate.Struct(input)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	err = h.authService.SignUp(domain.UserInfo{Username: input.Username, Password: input.Password, Role: input.Role})
	if err != nil {
		if errors.Is(err, errs.ErrAlreadyExists) {
			response.Conflict(w, "User with this username already exists")
			return
		}
		log.Println(err)
		response.InternalServerError(w)
		return
	}
	response.OKMessage(w, "You signed up successfully")
}

// SignIn docs
//
//	@Summary		Вход пользователей
//	@Tags			auth
//	@Description	Вход для всех пользователей по логину и паролю
//	@ID				auth-sign-in
//	@Accept			json
//	@Produce		json
//	@Param			input		body		dto.SignInInput	true	"Логин и пароль"
//	@Success		200			{object}	dto.SignInOutput
//	@Failure		400,401,404	{object}	response.Body
//	@Failure		500			{object}	response.Body
//	@Failure		default		{object}	response.Body
//	@Router			/auth/sign_in [post]
func (h *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	var input dto.SignInInput

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	err = h.validate.Struct(input)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	tokens, err := h.authService.SignIn(domain.UserInfo{Username: input.Username, Password: input.Password})

	if err != nil {
		if errors.Is(err, errs.ErrUserNotExists) {
			response.NotFound(w, "User with this username not exists")
			return
		}
		if errors.Is(err, errs.ErrInvalidPass) {
			response.Unauthorized(w)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ansBytes, err := json.Marshal(
		dto.SignInOutput{
			AccessToken:  tokens.AccessToken,
			RefreshToken: tokens.RefreshToken,
		},
	)
	if err != nil {
		response.InternalServerError(w)
		return
	}

	response.WriteResponse(w, http.StatusOK, ansBytes)
}

// RefreshTokens docs
//
//	@Summary		Обновление токенов
//	@Tags			auth
//	@Description	Обновление токенов
//	@ID				auth-refresh
//	@Accept			json
//	@Produce		json
//	@Param			input		body		dto.RefreshInput	true	"Рефреш токен"
//	@Success		200			{object}	domain.Tokens
//	@Failure		400,401,404	{object}	response.Body
//	@Failure		500			{object}	response.Body
//	@Failure		default		{object}	response.Body
//	@Router			/auth/refresh [post]
func (h *AuthHandler) RefreshTokens(w http.ResponseWriter, r *http.Request) {
	var input dto.RefreshInput

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	err = h.validate.Struct(input)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	tokens, err := h.authService.RefreshTokens(input.RefreshToken)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ansBytes, err := json.Marshal(
		dto.SignInOutput{
			AccessToken:  tokens.AccessToken,
			RefreshToken: tokens.RefreshToken,
		},
	)
	if err != nil {
		response.InternalServerError(w)
		return
	}

	response.WriteResponse(w, http.StatusOK, ansBytes)
}
