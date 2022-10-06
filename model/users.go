package model

type User struct {
	ID        int64  `json:"id" dynamodbav:"ID"`
	FirstName string `json:"first_name" dynamodbav:"Fn"`
}
