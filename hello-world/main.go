package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func hello_handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	name, ok := request.QueryStringParameters["name"]
	if !ok {
		log.Print("No name passed in, using default 'World'")
		name = "World"
	}

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("Hello %s", name),
		StatusCode: 200,
	}, nil
}

func goodbye_handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	name, ok := request.QueryStringParameters["name"]
	if !ok {
		log.Print("No name passed in, using default 'World'")
		name = "World"
	}

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("Goodbye %s", name),
		StatusCode: 200,
	}, nil
}

func main() {
	switch os.Getenv("FUNCTION") {
	case "Hello":
		lambda.Start(hello_handler)
	case "Goodbye":
		lambda.Start(goodbye_handler)
	}
}
