package registry

import (
	"github.com/ohunyk/webserver/interface/controller"
	"go.mongodb.org/mongo-driver/mongo"
)

type registry struct {
	db *mongo.Database
}

type Registry interface {
	NewAppController() controller.AppController
}

func NewRegistry(db *mongo.Database) Registry {
	return &registry{db}
}

func (r *registry) NewAppController() controller.AppController {
	return r.NewUserController()
}