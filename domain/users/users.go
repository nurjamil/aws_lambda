package users

import (
	"aws_lambda_go/model"
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type UsersDomainItf interface {
	CreateUser(ctx context.Context, user *model.User) (model.User, error)
}

type users struct {
	opt Options
	db  *dynamodb.Client
}

type Options struct {
	TableName string
}

func NewUsersDomainItf(opt Options, db *dynamodb.Client) UsersDomainItf {
	return &users{
		opt: opt,
		db:  db,
	}
}
