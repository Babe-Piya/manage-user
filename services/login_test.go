package services

import (
	"context"
	"testing"

	"manage-user/appconstants"
	mockrepositories "manage-user/mocks/repositories"
	"manage-user/repositories"

	"github.com/stretchr/testify/assert"
)

const (
	mockEncryptPassword = "mock-encrypt-password"
)

func TestLoginWhenGetUserErrorShouldReturnError(t *testing.T) {
	mockRepo := mockrepositories.NewMockUserRepository(t)
	mockGetUser := repositories.User{
		Email: mockEmail,
	}
	mockRepo.EXPECT().GetUserByFilter(context.TODO(), mockGetUser).
		Return(nil, mockError)

	service := NewUserService(mockRepo, nil)

	mockReq := LoginRequest{
		Email:    mockEmail,
		Password: mockPassword,
	}
	actual, err := service.Login(context.TODO(), mockReq)

	assert.Nil(t, actual)
	assert.ErrorIs(t, mockError, err)
}

func TestLoginWhenPasswordNotCorrectShouldReturnError(t *testing.T) {
	mockRepo := mockrepositories.NewMockUserRepository(t)
	mockGetUser := repositories.User{
		Email: mockEmail,
	}
	mockGetUserResp := []repositories.User{{
		ID:       mockID,
		Email:    mockEmail,
		Name:     mockName,
		Password: mockEncryptPassword,
	}}
	mockRepo.EXPECT().GetUserByFilter(context.TODO(), mockGetUser).
		Return(mockGetUserResp, nil)

	service := NewUserService(mockRepo, nil)

	mockReq := LoginRequest{
		Email:    mockEmail,
		Password: mockPassword,
	}
	actual, err := service.Login(context.TODO(), mockReq)

	assert.Nil(t, actual)
	assert.ErrorIs(t, appconstants.WrongKeyLoginError, err)
}
