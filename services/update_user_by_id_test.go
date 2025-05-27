package services

import (
	"context"
	"errors"
	"testing"

	"manage-user/appconstants"
	mockrepositories "manage-user/mocks/repositories"
	"manage-user/repositories"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

const (
	mockID       = "mock-id"
	mockEmail    = "mock-email"
	mockName     = "mock-name"
	mockOldEmail = "mock-old-email"
)

var (
	mockError = errors.New("mock error")
)

func TestUpdateUserByIDWhenNotErrorShouldReturnSuccess(t *testing.T) {
	mockRepo := mockrepositories.NewMockUserRepository(t)
	mockGetUser := repositories.User{
		ID:    mockID,
		Email: mockEmail,
	}
	mockGetUserResp := []repositories.User{{ID: mockID, Email: mockOldEmail}}
	mockRepo.EXPECT().GetUserByFilter(context.TODO(), mockGetUser).
		Return(mockGetUserResp, nil)

	mockUpdate := repositories.User{
		ID:    mockID,
		Name:  mockName,
		Email: mockEmail,
	}
	mockRepo.EXPECT().UpdateUserByID(context.TODO(), mockUpdate).Return(nil)

	logger, _ := zap.NewProduction()
	service := NewUserService(mockRepo, nil, logger)

	mockReq := UpdateUserRequest{
		ID:    mockID,
		Name:  mockName,
		Email: mockEmail,
	}
	expected := &UpdateUserResponse{
		Code:    appconstants.SuccessCode,
		Message: appconstants.SuccessMessage,
	}
	actual, err := service.UpdateUserByID(context.TODO(), mockReq)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestUpdateUserByIDWhenIDNotFoundShouldReturnError(t *testing.T) {
	mockGetUser := repositories.User{
		ID:    mockID,
		Email: mockEmail,
	}
	mockRepo := mockrepositories.NewMockUserRepository(t)
	mockRepo.EXPECT().GetUserByFilter(context.TODO(), mockGetUser).
		Return(nil, mockError)

	logger, _ := zap.NewProduction()
	service := NewUserService(mockRepo, nil, logger)

	mockReq := UpdateUserRequest{
		ID:    mockID,
		Name:  mockName,
		Email: mockEmail,
	}

	actual, err := service.UpdateUserByID(context.TODO(), mockReq)

	assert.Nil(t, actual)
	assert.ErrorIs(t, mockError, err)
}

func TestUpdateUserByIDWhenEmailDuplicateShouldReturnError(t *testing.T) {
	mockGetUser := repositories.User{
		ID:    mockID,
		Email: mockEmail,
	}
	mockGetUserResp := []repositories.User{{ID: mockID, Email: mockEmail}}
	mockRepo := mockrepositories.NewMockUserRepository(t)
	mockRepo.EXPECT().GetUserByFilter(context.TODO(), mockGetUser).
		Return(mockGetUserResp, nil)

	logger, _ := zap.NewProduction()
	service := NewUserService(mockRepo, nil, logger)

	mockReq := UpdateUserRequest{
		ID:    mockID,
		Name:  mockName,
		Email: mockEmail,
	}

	actual, err := service.UpdateUserByID(context.TODO(), mockReq)

	assert.Nil(t, actual)
	assert.ErrorIs(t, appconstants.DuplicateEmailError, err)
}

func TestUpdateUserByIDWhenUpdateFailShouldReturnError(t *testing.T) {
	mockGetUser := repositories.User{
		ID:    mockID,
		Email: mockEmail,
	}
	mockGetUserResp := []repositories.User{{ID: mockID, Email: mockOldEmail}}
	mockRepo := mockrepositories.NewMockUserRepository(t)
	mockRepo.EXPECT().GetUserByFilter(context.TODO(), mockGetUser).
		Return(mockGetUserResp, nil)

	mockUpdate := repositories.User{
		ID:    mockID,
		Name:  mockName,
		Email: mockEmail,
	}
	mockRepo.EXPECT().UpdateUserByID(context.TODO(), mockUpdate).Return(mockError)

	logger, _ := zap.NewProduction()
	service := NewUserService(mockRepo, nil, logger)

	mockReq := UpdateUserRequest{
		ID:    mockID,
		Name:  mockName,
		Email: mockEmail,
	}
	actual, err := service.UpdateUserByID(context.TODO(), mockReq)

	assert.Nil(t, actual)
	assert.ErrorIs(t, mockError, err)
}
