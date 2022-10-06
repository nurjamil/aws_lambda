package main

import (
	"context"
	"os"

	"aws_lambda_go/domain"
	"aws_lambda_go/domain/users"
	"aws_lambda_go/handler"
	"aws_lambda_go/usecase"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
)

func main() {
	region := os.Getenv("AWS_REGION")
	cfg, err := config.LoadDefaultConfig(context.TODO(), func(o *config.LoadOptions) error {
		o.Region = region
		return nil
	})
	if err != nil {
		panic(err)
	}
	dynaClient := dynamodb.NewFromConfig(cfg)

	dynaClient.CreateTable(context.Background(), &dynamodb.CreateTableInput{
		TableName: aws.String("users"),
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("ID"),
				AttributeType: types.ScalarAttributeTypeN,
			},
			{
				AttributeName: aws.String("Fn"),
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("ID"),
				KeyType:       types.KeyTypeHash,
			},
		},
	})

	dom := domain.NewDomain(domain.Options{
		UsersOption: users.Options{
			TableName: "users",
		},
	}, dynaClient)

	usecase := usecase.NewUsecase(usecase.Options{}, dom)

	handlerItf := handler.NewHandler(handler.Options{}, usecase)

	lambda.Start(handlerItf.HandlerHelper)
}
