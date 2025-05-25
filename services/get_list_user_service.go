package services

import (
	"context"
	"log"
)

type GetListUserResponse struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Users   []UserDetail `json:"users"`
}

type UserDetail struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (srv *userService) GetListUser(ctx context.Context) (*GetListUserResponse, error) {
	users, err := srv.UserRepo.GetListUser(ctx)
	if err != nil {
		log.Println(err)

		return nil, err
	}

	var userDetails []UserDetail
	for _, user := range users {
		userDetails = append(userDetails, UserDetail{
			Name:  user.Name,
			Email: user.Email,
		})
	}

	return &GetListUserResponse{
		Code:    0,
		Message: "Success",
		Users:   userDetails,
	}, nil
}
