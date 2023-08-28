package inmemory

import (
	"context"
	"ddd-golang/domain"
	"sync"
)

type inMemoryUserRepository struct {
	storage map[string]domain.User
	mutex   sync.RWMutex
}

func Init(storage map[string]domain.User) domain.UserRepository {
	return &inMemoryUserRepository{
		storage: storage,
	}
}

func (u *inMemoryUserRepository) CreateUser(ctx context.Context, user domain.User) domain.User {
	u.mutex.Lock()
	defer u.mutex.Unlock()
	u.storage[user.EmailId] = user
	return user
}

func (u *inMemoryUserRepository) GetUserByEmailId(ctx context.Context, emailId string) (domain.User, error) {
	u.mutex.RLock()
	defer u.mutex.RUnlock()
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
