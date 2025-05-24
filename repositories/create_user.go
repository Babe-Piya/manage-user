package repositories

import (
	"context"
)

func (repo *userRepository) CreateUser(ctx context.Context, user User) (interface{}, error) {
	result, err := repo.DB.Collection("users").InsertOne(ctx, user)
	if err != nil {
		return "", err
	}

	return result.InsertedID, nil
}
