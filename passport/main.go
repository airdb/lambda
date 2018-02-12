package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Received body: ", request.Body)

	msg := `{"token":"416c9006ca431f503160b7e7","status":0}`
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    map[string]string{"x-airdb-token": "416c9006ca431f503160b7e7"},
		Body:       msg,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
