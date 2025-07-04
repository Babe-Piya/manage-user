package services

import (
	"context"

	"manage-user/appconfig"
	"manage-user/repositories"

	"go.uber.org/zap"
)

type UserService interface {
	CreateUser(ctx context.Context, req CreateUserRequest) (*CreateUserResponse, error)
	GetListUser(ctx context.Context) (*GetListUserResponse, error)
	GetUserByID(ctx context.Context, id string) (*GetUserResponse, error)
	UpdateUserByID(ctx context.Context, req UpdateUserRequest) (*UpdateUserResponse, error)
	DeleteUserByID(ctx context.Context, id string) (*DeleteUserResponse, error)
	Login(ctx context.Context, req LoginRequest) (*LoginResponse, error)
}

type userService struct {
	UserRepo repositories.UserRepository
	Config   *appconfig.AppConfig
	Log      *zap.Logger
}

func NewUserService(userRepo repositories.UserRepository, config *appconfig.AppConfig, log *zap.Logger) UserService {
	return &userService{
		UserRepo: userRepo,
		Config:   config,
		Log:      log,
	}
}
