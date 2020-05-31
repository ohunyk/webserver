package registry

import (
	"github.com/ohunyk/webserver/interface/controller"
	"github.com/ohunyk/webserver/interface/presenter"
	"github.com/ohunyk/webserver/interface/repository"
	"github.com/ohunyk/webserver/usecase/interactor"
)

func (r *registry) NewUserController() controller.UserController {
	return controller.NewUserController(r.NewUserInteractor())
}

func (r *registry) NewUserInteractor() interactor.UserInteractor {
	return interactor.NewUserInteractor(r.NewUserRepository(), r.NewUserPresenter())
}

func (r *registry) NewUserRepository() repository.UserRepository {
	return repository.NewRepository(r.db)
}

func (r *registry) NewUserPresenter() presenter.UserPresenter {
	return presenter.NewUserPresenter()
}