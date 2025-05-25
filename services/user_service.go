package services

import (
	"context"

	"manage-user/repositories"
)

type UserService interface {
	CreateUser(ctx context.Context, req CreateUserRequest) (*CreateUserResponse, error)
	GetListUser(ctx context.Context) (*GetListUserResponse, error)
	GetUserByID(ctx context.Context, id string) (*GetUserResponse, error)
	UpdateUserByID(ctx context.Context, req UpdateUserRequest) (*UpdateUserResponse, error)
	DeleteUserByID(ctx context.Context, id string) (*DeleteUserResponse, error)
}

type userService struct {
	UserRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{
		UserRepo: userRepo,
	}
}
