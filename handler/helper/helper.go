package helper

import (
	"log"

	"github.com/pkg/errors"

	"github.com/aws/aws-lambda-go/events"
	jsoniter "github.com/json-iterator/go"
)

func ApiResponse(status int, body interface{}) (*events.APIGatewayV2HTTPResponse, error) {
	resp := events.APIGatewayV2HTTPResponse{Headers: map[string]string{"Content-Type": "application/json"}}
	resp.StatusCode = status

	raw, err := jsoniter.Marshal(body)
	if err != nil {
		log.Fatalf("%+v", errors.Wrap(err, `fatal`))
	}

	resp.Body = string(raw)

	return &resp, nil
}
