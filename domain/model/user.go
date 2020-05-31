package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Email    string             `bson:"email,omitempty"`
	Password string             `bson:"password,omitempty"`
	Nickname string             `bson:"nickname,omitempty"`
}

//type Repository interface {
//	GetByEmail(ctx context.Context, email string) (User, error)
//}
//
//type Service interface {
//	GetByEmail(ctx context.Context, email string) (User, error)
//}
