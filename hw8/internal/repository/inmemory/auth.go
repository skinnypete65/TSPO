package inmemory

import (
	"ecom/internal/domain"
	"ecom/internal/errs"
	"ecom/internal/repository"
)

type AuthRepoInMemory struct {
	users map[string]domain.UserInfo
}

func NewAuthRepoInMemory() repository.AuthRepo {
	return &AuthRepoInMemory{
		users: make(map[string]domain.UserInfo),
	}
}

func (r *AuthRepoInMemory) CheckUserExists(uuid string) bool {
	_, exists := r.users[uuid]
	return exists
}

func (r *AuthRepoInMemory) InsertUser(user domain.UserInfo) error {
	r.users[user.ID] = user
	return nil
}
func (r *AuthRepoInMemory) GetUserByUserName(username string) (domain.UserInfo, error) {
	for _, userInfo := range r.users {
		if userInfo.Username == username {
			return userInfo, nil
		}
	}

	return domain.UserInfo{}, errs.ErrUserNotExists
}
