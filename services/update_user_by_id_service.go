package services

import (
	"context"
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
	err := srv.UserRepo.UpdateUserByID(ctx, repositories.User{
		ID:    req.ID,
		Name:  req.Name,
		Email: req.Email,
	})
	if err != nil {
		log.Println(err)

		return nil, err
	}

	return &UpdateUserResponse{
		Code:    0,
		Message: "Success",
	}, nil
}
