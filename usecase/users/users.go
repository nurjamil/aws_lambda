package users

import (
	userdomain "aws_lambda_go/domain/users"
	"aws_lambda_go/model"

	"context"
)

var User UserItf

type UserItf interface {
	CreateUser(ctx context.Context, user *model.User) (model.User, error)
}

type users struct {
	opt         Options
	usersDomain userdomain.UsersDomainItf
}

type Options struct {
	TableName string
}

func NewUsers(opt Options, usersDomain userdomain.UsersDomainItf) UserItf {
	return &users{
		opt:         opt,
		usersDomain: usersDomain,
	}
}
