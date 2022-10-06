package users

import (
	"aws_lambda_go/model"
	"context"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/pkg/errors"
)

func (u *users) CreateUser(ctx context.Context, user *model.User) (model.User, error) {
	var result model.User
	output, err := u.db.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: &u.opt.TableName,
		Item: map[string]types.AttributeValue{
			"ID": &types.AttributeValueMemberN{Value: strconv.FormatInt(user.ID, 10)},
			"Fn": &types.AttributeValueMemberS{Value: user.FirstName},
		},
		ReturnValues: types.ReturnValueAllNew,
	})

	if err != nil {
		return model.User{}, errors.Wrap(err, `failed when put item in dynamodb`)
	}

	if err := attributevalue.UnmarshalMap(output.Attributes, &result); err != nil {
		return model.User{}, errors.Wrap(err, `failed when unmarshal the result`)
	}

	return result, nil

}
