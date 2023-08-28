package domain

import "context"

// User Domain
type User struct {
	Name    string `json:"name,omitempty"`
	EmailId string `json:"emailId,omitempty"`
}

// UserRepository contains meaningful method names but without any business logic
type UserRepository interface {
	// CreateUser creates user to persistent layer(i.e database)
	CreateUser(ctx context.Context, user User) User
	GetUserByEmailId(ctx context.Context, emailId string) (User, error)
	CheckEmailId(ctx context.Context, emailId string) bool
}

// UserUseCase contains meaningful method names which would contain business logic
type UserUseCase interface {
	// CreateUser during concrete implementation, creates user as per business logic i.e. check duplicate emailId etc
	// and then uses user repository to create user
	CreateUser(ctx context.Context, user User) (User, error)
	GetUserByEmailId(ctx context.Context, emailId string) (User, error)
}
