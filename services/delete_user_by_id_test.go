package services

import (
	"context"
	"testing"

	"manage-user/appconstants"
	mockrepositories "manage-user/mocks/repositories"
	"manage-user/repositories"

	"github.com/stretchr/testify/assert"
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

	service := NewUserService(mockRepo, nil)

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

	service := NewUserService(mockRepo, nil)

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

	service := NewUserService(mockRepo, nil)

	actual, err := service.DeleteUserByID(context.TODO(), mockID)

	assert.Nil(t, actual)
	assert.ErrorIs(t, mockError, err)
}
