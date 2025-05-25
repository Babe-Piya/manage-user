package services

import (
	"context"
	"log"
)

type DeleteUserResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (srv *userService) DeleteUserByID(ctx context.Context, id string) (*DeleteUserResponse, error) {
	err := srv.UserRepo.DeleteUserByID(ctx, id)
	if err != nil {
		log.Println(err)

		return nil, err
	}

	return &DeleteUserResponse{
		Code:    0,
		Message: "Success",
	}, nil
}
