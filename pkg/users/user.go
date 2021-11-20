package users

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)



type User struct{
	Email     string `json:"email"`
	FirstName  string `json:"firstName"`
	LastName   string  `json:"lastName"`
}


func FetchUser(email string , tablename string , dynaClient dynamodbiface.DynamoDBAPI) (*User, error) {
	input := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"email":{
				S: aws.String(email),
			}
		},
		TableName: aws.String(tablename),
	}
	result , err := dynaClient.GetItem(input)

	if err != nil {
		return nil , errors.New(ErrorFailedToFetchRecord)
	}
	item := new(User)
	err = dynamodbattribute.UnmarshalMap(result.Item, item)
	if err != nil {
		return nil , errors.New(ErrorFailedToUnMashalRecord)
	}
	return item , nil 
}

func FetchUsers(tablename string , dynaClient dynamodbiface.DynamoDBAPI)(*[]User,error)  {
	input := &dynamodb.ScanInput{
		TableName: aws.String(tablename),
	}
	result, err := dynaClient.Scan(input)
	if err != nil {
		return nil , errors.New(ErrorFailedToFetchRecord)
	}
	item := new([]user)
	err = dynamodbattribute.UnmarshalMap(result.Items, item)

}
func CreateUser()  {
	
}
func UpdateUser()  {
	
}
func DeleteUser() error {
	
}