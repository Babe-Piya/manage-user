package repositories

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func (repo *userRepository) GetUserByFilter(ctx context.Context, filter User) ([]User, error) {
	var fil bson.M
	if filter.Name != "" {
		fil = bson.M{"name": filter.Name}
	}
	if filter.Email != "" {
		fil = bson.M{"email": filter.Email}
	}
	if filter.ID != "" {
		fil = bson.M{"_id": filter.ID}
	}

	cursor, err := repo.DB.Collection("users").Find(ctx, fil)
	if err != nil {
		return nil, err
	}

	var result []User
	if err = cursor.All(ctx, &result); err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, errors.New("user not found")
	}

	return result, nil
}
