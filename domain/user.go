package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Email  string             `bson:"email,omitempty"`
	Password string             `bson:"password,omitempty"`
	Nickname   string           `bson:"nickname,omitempty"`
}

type UserRepository interface {
	GetByEmail(ctx context.Context, email string) (User, error)
}

type UserService interface {
	GetByEmail(ctx context.Context, email string) (User, error)
}