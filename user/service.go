package user

import (
	"context"
	"github.com/ohunyk/webserver/domain"
	"time"
)

type userService struct {
	userRepo domain.UserRepository
	contextTimeout time.Duration
}

func NewUserService(ur domain.UserRepository, timeout time.Duration) domain.UserService{
	return &userService{
		userRepo: ur,
		contextTimeout: timeout,
	}
}

func (us *userService) GetByEmail(ctx context.Context, email string) (res domain.User, err error) {
	ctx, cancel := context.WithTimeout(ctx, us.contextTimeout)
	defer cancel()
	res, err = us.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return
	}
	return
}