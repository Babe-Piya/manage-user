package services

import (
	"context"
	"log"

	"manage-user/appconstants"
	"manage-user/repositories"
)

type GetUserResponse struct {
	Code    int        `json:"code"`
	Message string     `json:"message"`
	User    UserDetail `json:"user"`
}

func (srv *userService) GetUserByID(ctx context.Context, id string) (*GetUserResponse, error) {
	users, err := srv.UserRepo.GetUserByFilter(ctx, repositories.User{
		ID: id,
	})
	if err != nil {
		log.Println(err)

		return nil, err
	}

	return &GetUserResponse{
		Code:    appconstants.SuccessCode,
		Message: appconstants.SuccessMessage,
		User: UserDetail{
			ID:    users[0].ID,
			Name:  users[0].Name,
			Email: users[0].Email,
		},
	}, nil
}
