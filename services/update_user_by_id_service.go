package services

import (
	"context"

	"manage-user/appconstants"
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
	srv.Log.Info("Function UpdateUserByID")
	existingUser, err := srv.UserRepo.GetUserByFilter(ctx, repositories.User{ID: req.ID, Email: req.Email})
	if err != nil {
		srv.Log.Error(err.Error())

		return nil, err
	}

	for _, user := range existingUser {
		if user.Email == req.Email {
			srv.Log.Error("duplicate email")

			return nil, appconstants.DuplicateEmailError
		}
	}

	errUpdate := srv.UserRepo.UpdateUserByID(ctx, repositories.User{
		ID:    req.ID,
		Name:  req.Name,
		Email: req.Email,
	})
	if errUpdate != nil {
		srv.Log.Error(errUpdate.Error())

		return nil, errUpdate
	}

	return &UpdateUserResponse{
		Code:    appconstants.SuccessCode,
		Message: appconstants.SuccessMessage,
	}, nil
}
