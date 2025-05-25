package repositories

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func (repo *userRepository) DeleteUserByID(ctx context.Context, id string) error {
	_, err := repo.DB.Collection("users").DeleteOne(ctx, bson.D{{"_id", id}})
	if err != nil {
		return err
	}

	return nil
}
