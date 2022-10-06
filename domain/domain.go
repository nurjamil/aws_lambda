package domain

import (
	"aws_lambda_go/domain/users"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type Domain struct {
	Users users.UsersDomainItf
}

type Options struct {
	UsersOption users.Options
}

func NewDomain(
	opt Options,
	db *dynamodb.Client,
) Domain {
	return Domain{
		Users: users.NewUsersDomainItf(
			opt.UsersOption,
			db,
		),
	}
}
