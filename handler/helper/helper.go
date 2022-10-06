package helper

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/francoispqt/gojay"
)

func ApiResponse(status int, body interface{}) (*events.APIGatewayProxyResponse, error) {
	resp := events.APIGatewayProxyResponse{Headers: map[string]string{"Content-Type": "application/json"}}
	resp.StatusCode = status

	stringBody, _ := json.Marshal(body)
	resp.Body = string(stringBody)

	enc := gojay.BorrowEncoder(&strings.Builder{})
	defer enc.Release()
	if err := enc.Encode(body); err != nil {
		log.Fatal(err)
	}

	return &resp, nil
}
