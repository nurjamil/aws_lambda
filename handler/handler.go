package handler

import (
	"aws_lambda_go/usecase"

	"github.com/aws/aws-lambda-go/events"
)

type HandlerItf interface {
	CreateUser(req events.APIGatewayV2HTTPRequest) (*events.APIGatewayV2HTTPResponse,
		error)
	HandlerHelper(req events.APIGatewayV2HTTPRequest) (*events.APIGatewayV2HTTPResponse, error)
}

type handler struct {
	opt     Options
	usecase usecase.Usecase
}

type Options struct{}

func NewHandler(opt Options,
	usecase usecase.Usecase) HandlerItf {
	return &handler{
		opt:     opt,
		usecase: usecase,
	}
}
