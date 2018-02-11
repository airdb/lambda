package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Received body: ", request.Body)

	msg := "Hello, aws lambda go!\nGithub: https://github.com/airdb/lambda\n"
	return events.APIGatewayProxyResponse{Body: msg, StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
