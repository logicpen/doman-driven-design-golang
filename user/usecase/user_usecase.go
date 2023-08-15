package usecase

import (
	"context"
	"ddd-golang/domain"
	"errors"
)

type UserUseCase struct {
	userRepository domain.UserRepository
}

func NewUserUseCase(u domain.UserRepository) domain.UserUseCase {
	return &UserUseCase{
		userRepository: u,
	}
}

// CreateUser creates user as per business logic i.e check duplicate emailId etc
// and then uses user repository to create user
func (u *UserUseCase) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {
	if u.userRepository.CheckEmailId(ctx, user.EmailId) {
		return domain.User{}, errors.New("email id already registered")
	}
	createdUser := u.userRepository.CreateUser(ctx, user)
	return createdUser, nil
}

func (u *UserUseCase) GetUserByEmailId(ctx context.Context, emailId string) (domain.User, error) {
	user, err := u.userRepository.GetUserByEmailId(ctx, emailId)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}
