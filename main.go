package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/colinlcrawford/alexa-whiskey-skill/alexa"
	"github.com/colinlcrawford/alexa-whiskey-skill/whiskeyapi"
)

func HandleRequest() (alexa.Response, error) {
	randomFact := whiskeyapi.GetRandomWhiskeyFact()
	return alexa.NewResponse(randomFact), nil
}

func main() {
	lambda.Start(HandleRequest)
}
