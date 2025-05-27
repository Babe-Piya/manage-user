package services

import (
	"context"
	"manage-user/appconstants"
	"manage-user/repositories"
)

type DeleteUserResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (srv *userService) DeleteUserByID(ctx context.Context, id string) (*DeleteUserResponse, error) {
	srv.Log.Info("function DeleteUserByID")
	_, err := srv.UserRepo.GetUserByFilter(ctx, repositories.User{ID: id})
	if err != nil {
		srv.Log.Error(err.Error())

		return nil, err
	}

	errDel := srv.UserRepo.DeleteUserByID(ctx, id)
	if errDel != nil {
		srv.Log.Error(errDel.Error())

		return nil, errDel
	}

	return &DeleteUserResponse{
		Code:    appconstants.SuccessCode,
		Message: appconstants.SuccessMessage,
	}, nil
}
