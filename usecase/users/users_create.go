package users

import (
	"aws_lambda_go/model"
	"context"
)

func (u *users) CreateUser(ctx context.Context, user *model.User) (model.User, error) {
	return u.usersDomain.CreateUser(ctx, user)
}
