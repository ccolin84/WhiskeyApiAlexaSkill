package main

import (
    "github.com/ccolin84/alexa-whiskey-skill/whiskeyapi"
    "github.com/ccolin84/alexa-whiskey-skill/alexa"
    "github.com/aws/aws-lambda-go/lambda"
)


func HandleRequest() (alexa.Response, error) {
    randomFact := whiskeyapi.GetRandomWhiskeyFact()
    return alexa.NewResponse(randomFact), nil
}

func main() {
    lambda.Start(HandleRequest)
}
