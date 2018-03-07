package main

import (
	"encoding/json"
	"fmt"
	// "github.com/airdb/lambda/test/models"
	"./models"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Received body: ", request.Body)

	session := &models.Session{}
	session.Username = "dean"
	session.Token = "416c9006ca431f503160b7e7"
	session.Mail = "dean@airdb.com"
	msg := `{"token":"416c9006ca431f503160b7e7","status":0}`
	aa, _ := json.Marshal(session)
	msg = string(aa)
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    map[string]string{"x-airdb-token": "416c9006ca431f503160b7e7", "Access-Control-Allow-Origin": "*", "Cookie": "sessionid=982EF67276E9F8AF8B2CDA33B0C04DE1; username=dean; token=416c9006ca431f503160b7e7"},
		Body:       msg,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
