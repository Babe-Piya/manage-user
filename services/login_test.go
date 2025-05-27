package services

import (
	"context"
	"testing"

	"manage-user/appconstants"
	mockrepositories "manage-user/mocks/repositories"
	"manage-user/repositories"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
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

	logger, _ := zap.NewProduction()
	service := NewUserService(mockRepo, nil, logger)

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

	logger, _ := zap.NewProduction()
	service := NewUserService(mockRepo, nil, logger)

	mockReq := LoginRequest{
		Email:    mockEmail,
		Password: mockPassword,
	}
	actual, err := service.Login(context.TODO(), mockReq)

	assert.Nil(t, actual)
	assert.ErrorIs(t, appconstants.WrongKeyLoginError, err)
}
