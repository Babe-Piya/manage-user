package services

import (
	"context"
	"testing"

	"manage-user/appconstants"
	mockrepositories "manage-user/mocks/repositories"
	"manage-user/repositories"

	"github.com/stretchr/testify/assert"
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

	service := NewUserService(mockRepo, nil)

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

	service := NewUserService(mockRepo, nil)

	actual, err := service.GetUserByID(context.TODO(), mockID)

	assert.Nil(t, actual)
	assert.ErrorIs(t, mockError, err)
}
