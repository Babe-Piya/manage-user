package services

import (
	"context"
	"strings"
	"time"

	"manage-user/appconstants"
	"manage-user/repositories"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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
	srv.Log.Info("function CreateUser")
	existingUser, err := srv.UserRepo.GetUserByFilter(ctx, repositories.User{Email: req.Email})
	if err != nil && !strings.Contains(err.Error(), "not found") {
		srv.Log.Error(err.Error())

		return nil, err
	}

	if len(existingUser) != 0 {
		srv.Log.Error("duplicate email")

		return nil, appconstants.DuplicateEmailError
	}

	bytesPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		srv.Log.Error(err.Error())

		return nil, err
	}

	id, err := srv.UserRepo.CreateUser(ctx, repositories.User{
		ID:        uuid.New().String(),
		Name:      req.Name,
		Email:     req.Email,
		Password:  string(bytesPass),
		CreatedAt: time.Now(),
	})
	if err != nil {
		srv.Log.Error(err.Error())

		return nil, err
	}

	return &CreateUserResponse{
		Code:    appconstants.SuccessCode,
		Message: appconstants.SuccessMessage,
		ID:      id,
	}, nil
}
