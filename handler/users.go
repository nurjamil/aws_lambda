package handler

import (
	"context"
	"fmt"
	"net/http"

	"aws_lambda_go/handler/helper"
	"aws_lambda_go/handler/response"
	"aws_lambda_go/model"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func (h *handler) CreateUser(req events.APIGatewayV2HTTPRequest) (
	*events.APIGatewayV2HTTPResponse,
	error,
) {
	var user model.User

	// decode the request to model
	if err := jsoniter.Unmarshal([]byte(req.Body), &user); err != nil {
		helper.ApiResponse(http.StatusBadRequest, response.ResponseError{
			ErrorMsg: aws.String(errors.Wrap(err, `failed to unmarshal`).Error()),
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

func (h *handler) HandlerHelper(req events.APIGatewayV2HTTPRequest) (*events.APIGatewayV2HTTPResponse, error) {
	log.WithField(`request`, fmt.Sprintf("%+v", req)).Info(`request`)

	switch req.RequestContext.HTTP.Method {
	case "POST":
		return h.CreateUser(req)
	}

	return nil, nil

}
