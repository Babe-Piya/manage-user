package services

import (
	"context"
	"testing"

	"manage-user/appconstants"
	mockrepositories "manage-user/mocks/repositories"
	"manage-user/repositories"

	"github.com/stretchr/testify/assert"
)

func TestGetListUserWhenNotErrorShouldReturnSuccess(t *testing.T) {
	mockRepo := mockrepositories.NewMockUserRepository(t)
	mockGetUserResp := []repositories.User{{
		ID:    mockID,
		Email: mockEmail,
		Name:  mockName,
	}}
	mockRepo.EXPECT().GetListUser(context.TODO()).
		Return(mockGetUserResp, nil)

	service := NewUserService(mockRepo, nil)

	expected := &GetListUserResponse{
		Code:    appconstants.SuccessCode,
		Message: appconstants.SuccessMessage,
		Users: []UserDetail{{
			ID:    mockID,
			Email: mockEmail,
			Name:  mockName,
		}},
	}
	actual, err := service.GetListUser(context.TODO())

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func TestGetListUserWhenErrorShouldReturnError(t *testing.T) {
	mockRepo := mockrepositories.NewMockUserRepository(t)
	mockRepo.EXPECT().GetListUser(context.TODO()).
		Return(nil, mockError)

	service := NewUserService(mockRepo, nil)

	actual, err := service.GetListUser(context.TODO())

	assert.Nil(t, actual)
	assert.ErrorIs(t, mockError, err)
}
