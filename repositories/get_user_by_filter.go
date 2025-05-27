package repositories

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func (repo *userRepository) GetUserByFilter(ctx context.Context, filter User) ([]User, error) {
	var condition []bson.M
	if filter.Name != "" {
		condition = append(condition, bson.M{"name": filter.Name})
	}
	if filter.Email != "" {
		condition = append(condition, bson.M{"email": filter.Email})
	}
	if filter.ID != "" {
		condition = append(condition, bson.M{"_id": filter.ID})
	}

	fil := bson.M{"$or": condition}

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
