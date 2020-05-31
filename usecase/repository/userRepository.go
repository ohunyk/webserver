package repository

import (
	"github.com/ohunyk/webserver/domain/model"
)

type UserRepository interface {
	FindAll(u []*model.User) ([]*model.User, error)
}
