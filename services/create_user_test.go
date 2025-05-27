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
	mockPassword = "mock-password"
)

func TestCreateUserWhenGetUserErrorShouldReturnError(t *testing.T) {
	mockRepo := mockrepositories.NewMockUserRepository(t)
	mockGetUser := repositories.User{
		Email: mockEmail,
	}
	mockRepo.EXPECT().GetUserByFilter(context.TODO(), mockGetUser).
		Return(nil, mockError)

	logger, _ := zap.NewProduction()
	service := NewUserService(mockRepo, nil, logger)

	mockReq := CreateUserRequest{
		Name:     mockName,
		Email:    mockEmail,
		Password: mockPassword,
	}
	actual, err := service.CreateUser(context.TODO(), mockReq)

	assert.Nil(t, actual)
	assert.ErrorIs(t, mockError, err)
}

func TestCreateUserWhenDuplicateEmailShouldReturnError(t *testing.T) {
	mockRepo := mockrepositories.NewMockUserRepository(t)
	mockGetUser := repositories.User{
		Email: mockEmail,
	}
	mockGetUserResp := []repositories.User{{ID: mockID, Email: mockEmail}}
	mockRepo.EXPECT().GetUserByFilter(context.TODO(), mockGetUser).
		Return(mockGetUserResp, nil)

	logger, _ := zap.NewProduction()
	service := NewUserService(mockRepo, nil, logger)

	mockReq := CreateUserRequest{
		Name:     mockName,
		Email:    mockEmail,
		Password: mockPassword,
	}
	actual, err := service.CreateUser(context.TODO(), mockReq)

	assert.Nil(t, actual)
	assert.ErrorIs(t, appconstants.DuplicateEmailError, err)
}
