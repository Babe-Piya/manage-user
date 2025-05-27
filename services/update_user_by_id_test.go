package services

import (
	"context"
	"errors"
	"testing"

	"manage-user/appconstants"
	mockrepositories "manage-user/mocks/repositories"
	"manage-user/repositories"

	"github.com/stretchr/testify/assert"
)

const (
	mockID          = "mock-id"
	mockUpdateEmail = "mock-update-email"
	mockUpdateName  = "mock-update-name"
	mockOldEmail    = "mock-old-email"
)

var (
	mockError = errors.New("mock error")
)

func TestUpdateUserByIDWhenNotErrorShouldReturnSuccess(t *testing.T) {
	mockGetUser := repositories.User{
		ID:    mockID,
		Email: mockUpdateEmail,
	}
	mockGetUserResp := []repositories.User{{ID: mockID, Email: mockOldEmail}}
	mockRepo := mockrepositories.NewMockUserRepository(t)
	mockRepo.EXPECT().GetUserByFilter(context.TODO(), mockGetUser).
		Return(mockGetUserResp, nil)

	mockUpdate := repositories.User{
		ID:    mockID,
		Name:  mockUpdateName,
		Email: mockUpdateEmail,
	}
	mockRepo.EXPECT().UpdateUserByID(context.TODO(), mockUpdate).Return(nil)

	service := NewUserService(mockRepo, nil)

	mockReq := UpdateUserRequest{
		ID:    mockID,
		Name:  mockUpdateName,
		Email: mockUpdateEmail,
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
		Email: mockUpdateEmail,
	}
	mockRepo := mockrepositories.NewMockUserRepository(t)
	mockRepo.EXPECT().GetUserByFilter(context.TODO(), mockGetUser).
		Return(nil, mockError)

	service := NewUserService(mockRepo, nil)

	mockReq := UpdateUserRequest{
		ID:    mockID,
		Name:  mockUpdateName,
		Email: mockUpdateEmail,
	}

	actual, err := service.UpdateUserByID(context.TODO(), mockReq)

	assert.Nil(t, actual)
	assert.ErrorIs(t, mockError, err)
}

func TestUpdateUserByIDWhenEmailDuplicateShouldReturnError(t *testing.T) {
	mockGetUser := repositories.User{
		ID:    mockID,
		Email: mockUpdateEmail,
	}
	mockGetUserResp := []repositories.User{{ID: mockID, Email: mockUpdateEmail}}
	mockRepo := mockrepositories.NewMockUserRepository(t)
	mockRepo.EXPECT().GetUserByFilter(context.TODO(), mockGetUser).
		Return(mockGetUserResp, nil)

	service := NewUserService(mockRepo, nil)

	mockReq := UpdateUserRequest{
		ID:    mockID,
		Name:  mockUpdateName,
		Email: mockUpdateEmail,
	}

	actual, err := service.UpdateUserByID(context.TODO(), mockReq)

	assert.Nil(t, actual)
	assert.ErrorIs(t, duplicateEmailError, err)
}

func TestUpdateUserByIDWhenUpdateFailShouldReturnError(t *testing.T) {
	mockGetUser := repositories.User{
		ID:    mockID,
		Email: mockUpdateEmail,
	}
	mockGetUserResp := []repositories.User{{ID: mockID, Email: mockOldEmail}}
	mockRepo := mockrepositories.NewMockUserRepository(t)
	mockRepo.EXPECT().GetUserByFilter(context.TODO(), mockGetUser).
		Return(mockGetUserResp, nil)

	mockUpdate := repositories.User{
		ID:    mockID,
		Name:  mockUpdateName,
		Email: mockUpdateEmail,
	}
	mockRepo.EXPECT().UpdateUserByID(context.TODO(), mockUpdate).Return(mockError)

	service := NewUserService(mockRepo, nil)

	mockReq := UpdateUserRequest{
		ID:    mockID,
		Name:  mockUpdateName,
		Email: mockUpdateEmail,
	}
	actual, err := service.UpdateUserByID(context.TODO(), mockReq)

	assert.Nil(t, actual)
	assert.ErrorIs(t, mockError, err)
}
