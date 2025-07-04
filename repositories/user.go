package repositories

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user User) (interface{}, error)
	GetListUser(ctx context.Context) ([]User, error)
	GetUserByFilter(ctx context.Context, filter User) ([]User, error)
	UpdateUserByID(ctx context.Context, user User) error
	DeleteUserByID(ctx context.Context, id string) error
}

type userRepository struct {
	DB *mongo.Database
}

func NewUserRepository(db *mongo.Database) UserRepository {
	return &userRepository{db}
}

type User struct {
	ID        string    `json:"id" bson:"_id"`
	Name      string    `json:"name" bson:"name"`
	Email     string    `json:"email" bson:"email"`
	Password  string    `json:"password" bson:"password"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}
