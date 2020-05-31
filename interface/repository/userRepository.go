package repository

import (
	"context"
	"github.com/ohunyk/webserver/domain/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	Collection *mongo.Collection
}

type UserRepository interface {
	FindAll(u []*model.User) ([]*model.User, error)
}

func NewRepository(db *mongo.Database) UserRepository {
	return &userRepository{
		Collection: db.Collection("user"),
	}
}

func (ur *userRepository) FindAll(u []*model.User) ([]*model.User, error) {
	cursor, err := ur.Collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &u); err == nil {
		return u, err
	}
	return nil, err
}
