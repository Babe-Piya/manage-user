package services

import (
	"context"
	"errors"
	"log"
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
	existingUser, err := srv.UserRepo.GetUserByFilter(ctx, repositories.User{Email: req.Email})
	if err != nil && !strings.Contains(err.Error(), "not found") {
		log.Println(err)

		return nil, err
	}

	if len(existingUser) != 0 {
		log.Println("duplicate email")

		return nil, errors.New("error duplicate email")
	}

	bytesPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)

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
		log.Println(err)

		return nil, err
	}

	return &CreateUserResponse{
		Code:    appconstants.SuccessCode,
		Message: appconstants.SuccessMessage,
		ID:      id,
	}, nil
}
