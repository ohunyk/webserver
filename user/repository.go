package user

import (
	"context"
	"fmt"
	"github.com/ohunyk/webserver/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	Collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) domain.UserRepository{
	return &userRepository{
		Collection: db.Collection("user"),
	}
}

func (u *userRepository) GetByEmail(ctx context.Context, email string) (domain.User, error) {
	var users []domain.User
	cursor, err := u.Collection.Find(ctx, bson.M{})
	if err != nil {
		return domain.User{}, err
	}
	if err = cursor.All(ctx, &users); err != nil {
		return domain.User{}, err
	}
	fmt.Println(users)
	return users[0], err
}