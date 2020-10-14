package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest() (AlexaResponse, error) {
	randomFact := GetRandomWhiskeyFact()
	return NewAlexaResponse(randomFact), nil
}

func main() {
	lambda.Start(HandleRequest)
}
