package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var (
	TableName = "org"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	result, err := getOrg("jbestor-test", "dev01-useast1")
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 404,
		}, err
	}

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("The org %s has the domain %s", result["_name"], result["vanity_domain"]),
		StatusCode: 200,
	}, nil
}

func getOrg(org string, pod string) (map[string]*dynamodb.AttributeValue, error) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := dynamodb.New(sess)

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(TableName),
		Key: map[string]*dynamodb.AttributeValue{
			"_name": {
				S: aws.String(org),
			},
			"_pod": {
				S: aws.String(pod),
			},
		},
	})
	if err != nil {
		return nil, err
	}

	return result.Item, nil
}

func main() {
	lambda.Start(handler)
}
