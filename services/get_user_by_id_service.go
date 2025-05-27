package services

import (
	"context"

	"manage-user/appconstants"
	"manage-user/repositories"
)

type GetUserResponse struct {
	Code    int        `json:"code"`
	Message string     `json:"message"`
	User    UserDetail `json:"user"`
}

func (srv *userService) GetUserByID(ctx context.Context, id string) (*GetUserResponse, error) {
	srv.Log.Info("Function GetUserByID")
	users, err := srv.UserRepo.GetUserByFilter(ctx, repositories.User{
		ID: id,
	})
	if err != nil {
		srv.Log.Error(err.Error())

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
