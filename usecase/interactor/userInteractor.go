package interactor

import (
	"github.com/ohunyk/webserver/domain/model"
	"github.com/ohunyk/webserver/usecase/presenter"
	"github.com/ohunyk/webserver/usecase/repository"
)

type userInteractor struct {
	UserRepository repository.UserRepository
	UserPresenter presenter.UserPresenter
}

type UserInteractor interface {
	Get(u []*model.User) ([]*model.User, error)
}

func NewUserInteractor(r repository.UserRepository, p presenter.UserPresenter) UserInteractor {
	return &userInteractor{r, p}
}

func (ui *userInteractor) Get(u []*model.User) ([]*model.User, error) {
	u, err := ui.UserRepository.FindAll(u)
	if err != nil {
		return nil, err
	}
	return ui.UserPresenter.ResponseUsers(u), nil
}