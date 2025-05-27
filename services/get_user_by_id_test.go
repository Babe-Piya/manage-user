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

func TestGetUserByIDWhenNotErrorShouldReturnSuccess(t *testing.T) {
	mockRepo := mockrepositories.NewMockUserRepository(t)
	mockGetUser := repositories.User{
		ID: mockID,
	}
	mockGetUserResp := []repositories.User{{
		ID:    mockID,
		Email: mockEmail,
		Name:  mockName,
	}}
	mockRepo.EXPECT().GetUserByFilter(context.TODO(), mockGetUser).
		Return(mockGetUserResp, nil)

	logger, _ := zap.NewProduction()
	service := NewUserService(mockRepo, nil, logger)

	expected := &GetUserResponse{
		Code:    appconstants.SuccessCode,
		Message: appconstants.SuccessMessage,
		User: UserDetail{
			ID:    mockID,
			Email: mockEmail,
			Name:  mockName,
		},
	}
	actual, err := service.GetUserByID(context.TODO(), mockID)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestGetUserByIDWhenErrorShouldReturnError(t *testing.T) {
	mockRepo := mockrepositories.NewMockUserRepository(t)
	mockGetUser := repositories.User{
		ID: mockID,
	}
	mockRepo.EXPECT().GetUserByFilter(context.TODO(), mockGetUser).
		Return(nil, mockError)

	logger, _ := zap.NewProduction()
	service := NewUserService(mockRepo, nil, logger)

	actual, err := service.GetUserByID(context.TODO(), mockID)

	assert.Nil(t, actual)
	assert.ErrorIs(t, mockError, err)
}
