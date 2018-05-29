package main

import (
	"github.com/aws/aws-lamba-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Response struct {
	FirstName string	`json:"firstname"`
	LastName string		`json:"lastname"`
	AccountNum string	`json:"account"`
	ActiveLines []string `json:lines`
}

func Handler() (events.APIGatewayProxyResponse, error) {
	rsp := Response{
		FirstName: "First Name",
		LastName: "Last Name",
		AccountNum: "01234",
		ActiveLines: []string{"2134567890", "2134567891"},
	}
	return events.APIGatewayProxyResponse{Body: rsp, StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}

