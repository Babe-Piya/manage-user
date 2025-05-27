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

func TestDeleteUserByIDWhenNotErrorShouldReturnSuccess(t *testing.T) {
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

	mockRepo.EXPECT().DeleteUserByID(context.TODO(), mockID).Return(nil)

	logger, _ := zap.NewProduction()
	service := NewUserService(mockRepo, nil, logger)

	expected := &DeleteUserResponse{
		Code:    appconstants.SuccessCode,
		Message: appconstants.SuccessMessage,
	}
	actual, err := service.DeleteUserByID(context.TODO(), mockID)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestDeleteUserByIDWhenGetUserErrorShouldReturnError(t *testing.T) {
	mockRepo := mockrepositories.NewMockUserRepository(t)
	mockGetUser := repositories.User{
		ID: mockID,
	}
	mockRepo.EXPECT().GetUserByFilter(context.TODO(), mockGetUser).
		Return(nil, mockError)

	logger, _ := zap.NewProduction()
	service := NewUserService(mockRepo, nil, logger)

	actual, err := service.DeleteUserByID(context.TODO(), mockID)

	assert.Nil(t, actual)
	assert.ErrorIs(t, mockError, err)
}

func TestDeleteUserByIDWhenDeleteErrorShouldReturnError(t *testing.T) {
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

	mockRepo.EXPECT().DeleteUserByID(context.TODO(), mockID).Return(mockError)

	logger, _ := zap.NewProduction()
	service := NewUserService(mockRepo, nil, logger)

	actual, err := service.DeleteUserByID(context.TODO(), mockID)

	assert.Nil(t, actual)
	assert.ErrorIs(t, mockError, err)
}
