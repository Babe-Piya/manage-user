package services

import (
	"context"
	"errors"
	"log"

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
	existingUser, err := srv.UserRepo.GetUserByFilter(ctx, repositories.User{ID: req.ID, Email: req.Email})
	if err != nil {
		log.Println(err)

		return nil, err
	}

	for _, user := range existingUser {
		if user.Email == req.Email {
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
