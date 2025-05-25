package repositories

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func (repo *userRepository) GetListUser(ctx context.Context) ([]User, error) {
	cursor, err := repo.DB.Collection("users").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var result []User
	if err = cursor.All(ctx, &result); err != nil {
		return nil, err
	}

	return result, nil
}
