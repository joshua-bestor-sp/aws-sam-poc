package main

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestHandler(t *testing.T) {
	t.Run("Success with default name 'World'", func(t *testing.T) {
		response, err := handler(events.APIGatewayProxyRequest{})

		if response.Body != "Hello World" {
			t.Fatal("Should use the default name: 'World'")
		}
		if err != nil {
			t.Fatal("Everything should be ok")
		}
	})

	t.Run("Success with name passed in", func(t *testing.T) {
		response, err := handler(events.APIGatewayProxyRequest{
			QueryStringParameters: map[string]string{"name": "test"},
		})

		if response.Body != "Hello test" {
			t.Fatal("Should use the default name: 'World'")
		}
		if err != nil {
			t.Fatal("Everything should be ok")
		}
	})
}
