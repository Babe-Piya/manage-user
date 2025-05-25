package services

import (
	"context"
	"log"
	"time"

	"manage-user/repositories"

	"github.com/google/uuid"
)

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type CreateUserResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	ID      interface{} `json:"id"`
}

func (srv *userService) CreateUser(ctx context.Context, req CreateUserRequest) (*CreateUserResponse, error) {
	id, err := srv.UserRepo.CreateUser(ctx, repositories.User{
		ID:        uuid.New().String(),
		Name:      req.Name,
		Email:     req.Email,
		Password:  req.Password,
		CreatedAt: time.Now(),
	})
	if err != nil {
		log.Println(err)

		return nil, err
	}

	return &CreateUserResponse{
		Code:    0,
		Message: "Success",
		ID:      id,
	}, nil
}
