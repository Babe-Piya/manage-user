package services

import (
	"context"
	"errors"
	"log"
	"strings"

	"manage-user/repositories"
)

type UpdateUserRequest struct {
	ID    string `json:"id" validate:"required"`
	Name  string `json:"name"`
	Email string `json:"email" `
}

type UpdateUserResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (srv *userService) UpdateUserByID(ctx context.Context, req UpdateUserRequest) (*UpdateUserResponse, error) {
	if req.Email != "" {
		existingUser, err := srv.UserRepo.GetUserByFilter(ctx, repositories.User{Email: req.Email})
		if err != nil && !strings.Contains(err.Error(), "not found") {
			log.Println(err)

			return nil, err
		}

		if len(existingUser) != 0 {
			log.Println("duplicate email")

			return nil, errors.New("error duplicate email")
		}
	}

	errUpdate := srv.UserRepo.UpdateUserByID(ctx, repositories.User{
		ID:    req.ID,
		Name:  req.Name,
		Email: req.Email,
	})
	if errUpdate != nil {
		log.Println(errUpdate)

		return nil, errUpdate
	}

	return &UpdateUserResponse{
		Code:    0,
		Message: "Success",
	}, nil
}
