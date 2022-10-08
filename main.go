package main

import (
	"context"
	"fmt"

	"aws_lambda_go/domain"
	"aws_lambda_go/domain/users"
	"aws_lambda_go/handler"
	"aws_lambda_go/usecase"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func main() {
	// region := os.Getenv("AWS_REGION")
	cfg, err := config.LoadDefaultConfig(context.TODO()) // config.WithRegion("ca-central-1"),
	// config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
	// 	return aws.Endpoint{URL: "http://localhost:8000"}, nil
	// })),

	if err != nil {
		panic(err)
	}
	dynaClient := dynamodb.NewFromConfig(cfg)

	_, err = dynaClient.CreateTable(context.Background(), &dynamodb.CreateTableInput{
		TableName: aws.String("users"),
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("ID"),
				AttributeType: types.ScalarAttributeTypeN,
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("ID"),
				KeyType:       types.KeyTypeHash,
			},
		},
		BillingMode: types.BillingModePayPerRequest,
	})
	if err != nil {
		logrus.WithField(`warning`, fmt.Sprintf("%+v", errors.Wrap(err, `message`))).Warning("warning")
	}

	dom := domain.NewDomain(domain.Options{
		UsersOption: users.Options{
			TableName: "users",
		},
	}, dynaClient)

	usecase := usecase.NewUsecase(usecase.Options{}, dom)

	handlerItf := handler.NewHandler(handler.Options{}, usecase)

	lambda.Start(handlerItf.HandlerHelper)
}
