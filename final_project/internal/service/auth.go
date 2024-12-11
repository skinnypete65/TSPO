package service

import (
	"time"

	"ecom/internal/domain"
	"ecom/internal/errs"
	"ecom/internal/repository"
	"ecom/internal/tokens"
	"ecom/pkg/hash"
	"github.com/google/uuid"
)

type AuthService interface {
	SignUp(user domain.UserInfo) error
	SignIn(user domain.UserInfo) (domain.Tokens, error)
	RefreshTokens(refreshToken string) (domain.Tokens, error)
}

type authService struct {
	authRepo        repository.AuthRepo
	hasher          hash.PasswordHasher
	tokenManager    tokens.TokenManager
	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
}

func NewAuthService(
	authRepo repository.AuthRepo,
	hasher hash.PasswordHasher,
	tokenManager tokens.TokenManager,
) AuthService {
	return &authService{
		authRepo:        authRepo,
		hasher:          hasher,
		tokenManager:    tokenManager,
		accessTokenTTL:  2 * time.Hour,       // 2 hours for access
		refreshTokenTTL: 30 * 24 * time.Hour, // 30 days for refresh
	}
}

func (s *authService) SignUp(user domain.UserInfo) error {
	alreadyExists := s.authRepo.CheckUserExists(user.Username)
	if alreadyExists {
		return errs.ErrAlreadyExists
	}

	hashedPass, err := s.hasher.Hash(user.Password)
	if err != nil {
		return err
	}
	user.ID = uuid.New().String()
	user.Password = hashedPass

	err = s.authRepo.InsertUser(user)
	return err
}

func (s *authService) SignIn(info domain.UserInfo) (domain.Tokens, error) {
	user, err := s.authRepo.GetUserByUserName(info.Username)
	if err != nil {
		return domain.Tokens{}, err
	}

	inputHashPass, err := s.hasher.Hash(info.Password)
	if err != nil {
		return domain.Tokens{}, err
	}

	if user.Password != inputHashPass {
		return domain.Tokens{}, errs.ErrInvalidPass
	}

	accessToken, err := s.tokenManager.NewJWT(tokens.TokenInfo{
		UserID:   user.ID,
		UserRole: user.Role,
	}, s.accessTokenTTL)
	if err != nil {
		return domain.Tokens{}, err
	}

	refreshToken, err := s.tokenManager.NewJWT(tokens.TokenInfo{
		UserID:   user.ID,
		UserRole: user.Role,
	}, s.refreshTokenTTL)
	if err != nil {
		return domain.Tokens{}, err
	}

	return domain.Tokens{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}

func (s *authService) RefreshTokens(refreshToken string) (domain.Tokens, error) {
	tokenInfo, err := s.tokenManager.Parse(refreshToken)
	if err != nil {
		return domain.Tokens{}, err
	}

	accessToken, err := s.tokenManager.NewJWT(tokens.TokenInfo{
		UserID:   tokenInfo.UserID,
		UserRole: tokenInfo.UserRole,
	}, s.accessTokenTTL)
	if err != nil {
		return domain.Tokens{}, err
	}

	newRefreshToken, err := s.tokenManager.NewJWT(tokens.TokenInfo{
		UserID:   tokenInfo.UserID,
		UserRole: tokenInfo.UserRole,
	}, s.refreshTokenTTL)
	if err != nil {
		return domain.Tokens{}, err
	}

	return domain.Tokens{AccessToken: accessToken, RefreshToken: newRefreshToken}, nil
}
