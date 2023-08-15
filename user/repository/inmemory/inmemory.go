package inmemory

import (
	"context"
	"ddd-golang/domain"
)

type inMemoryUserRepository struct {
	storage map[string]domain.User
}

func Init(storage map[string]domain.User) domain.UserRepository {
	return &inMemoryUserRepository{
		storage: storage,
	}
}

func (u *inMemoryUserRepository) CreateUser(ctx context.Context, user domain.User) domain.User {
	u.storage[user.EmailId] = user
	return user
}

func (u *inMemoryUserRepository) GetUserByEmailId(ctx context.Context, emailId string) (domain.User, error) {
	return u.storage[emailId], nil
}

func (u *inMemoryUserRepository) CheckEmailId(ctx context.Context, emailId string) bool {
	for _, user := range u.storage {
		if user.EmailId == emailId {
			return true
		}
	}
	return false
}
