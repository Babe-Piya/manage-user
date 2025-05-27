package services

import (
	"context"
	"log"

	"manage-user/repositories"
)

type DeleteUserResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (srv *userService) DeleteUserByID(ctx context.Context, id string) (*DeleteUserResponse, error) {
	_, err := srv.UserRepo.GetUserByFilter(ctx, repositories.User{ID: id})
	if err != nil {
		log.Println(err)

		return nil, err
	}

	errDel := srv.UserRepo.DeleteUserByID(ctx, id)
	if errDel != nil {
		log.Println(errDel)

		return nil, errDel
	}

	return &DeleteUserResponse{
		Code:    0,
		Message: "Success",
	}, nil
}
