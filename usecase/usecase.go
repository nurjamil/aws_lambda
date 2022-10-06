package usecase

import (
	"aws_lambda_go/domain"
	"aws_lambda_go/usecase/users"
)

type Usecase struct {
	Users users.UserItf
}

type Options struct {
	UsersOptions users.Options
}

func NewUsecase(
	opt Options,
	dom domain.Domain,
) Usecase {
	return Usecase{
		Users: users.NewUsers(opt.UsersOptions,
			dom.Users),
	}
}
