package handler

import (
	"aws_lambda_go/usecase"

	"github.com/aws/aws-lambda-go/events"
)

type HandlerItf interface {
	CreateUser(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse,
		error)
	HandlerHelper(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error)
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
