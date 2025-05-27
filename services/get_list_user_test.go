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

func TestGetListUserWhenNotErrorShouldReturnSuccess(t *testing.T) {
	mockRepo := mockrepositories.NewMockUserRepository(t)
	mockGetUserResp := []repositories.User{{
		ID:    mockID,
		Email: mockEmail,
		Name:  mockName,
	}}
	mockRepo.EXPECT().GetListUser(context.TODO()).
		Return(mockGetUserResp, nil)

	logger, _ := zap.NewProduction()
	service := NewUserService(mockRepo, nil, logger)

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

	logger, _ := zap.NewProduction()
	service := NewUserService(mockRepo, nil, logger)

	actual, err := service.GetListUser(context.TODO())

	assert.Nil(t, actual)
	assert.ErrorIs(t, mockError, err)
}
