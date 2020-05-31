package presenter

import "github.com/ohunyk/webserver/domain/model"

type UserPresenter interface {
	ResponseUsers(u []*model.User) []*model.User
}
