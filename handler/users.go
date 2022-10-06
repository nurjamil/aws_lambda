package handler

import (
	"bytes"
	"context"
	"fmt"
	"net/http"

	"aws_lambda_go/handler/helper"
	"aws_lambda_go/handler/response"
	"aws_lambda_go/model"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/francoispqt/gojay"
)

func (h *handler) CreateUser(req events.APIGatewayProxyRequest) (
	*events.APIGatewayProxyResponse,
	error,
) {
	var user model.User

	// decode the request to model
	dec := gojay.BorrowDecoder(bytes.NewReader([]byte(req.Body)))
	defer dec.Release()
	if err := dec.Decode(&user); err != nil {
		helper.ApiResponse(http.StatusBadRequest, response.ResponseError{
			ErrorMsg: aws.String(err.Error()),
		})
	}

	result, err := h.usecase.Users.CreateUser(context.Background(), &user)
	if err != nil {
		return helper.ApiResponse(http.StatusBadRequest, response.ResponseError{
			ErrorMsg: aws.String(err.Error()),
			DebugMsg: fmt.Sprintf("%+v", err),
		})
	}

	return helper.ApiResponse(http.StatusCreated, result)
}

func (h *handler) HandlerHelper(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	switch req.HTTPMethod {
	case "POST":
		return h.CreateUser(req)
	}

	return nil, nil
}
