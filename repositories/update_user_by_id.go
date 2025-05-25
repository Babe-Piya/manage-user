package repositories

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func (repo *userRepository) UpdateUserByID(ctx context.Context, user User) error {
	updateFields := bson.M{}
	if user.Name != "" {
		updateFields["name"] = user.Name
	}
	if user.Email != "" {
		updateFields["email"] = user.Email
	}

	update := bson.M{
		"$set": updateFields,
	}
	_, err := repo.DB.Collection("users").UpdateOne(ctx, bson.D{{"_id", user.ID}}, update)
	if err != nil {
		return err
	}

	return nil
}
