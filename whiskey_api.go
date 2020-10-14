package main

import (
    "log"
    "net/http"
    "io/ioutil"
    "fmt"
    "encoding/json"
)

const WhiskeyApiUrl = "https://evening-citadel-85778.herokuapp.com:443"

type RandomFactApiResponse struct {
    Count int `json:"count"`
    Next int `json:"next"`
    Previous int `json:"previous"`
    Results []RandomFact `json:"results"`
}

type RandomFact struct {
    Id int `json:"id"`
    Text string `json:"text"`
    CreatedAt string `json:"created_at"`
}

func GetRandomWhiskeyFact() string {
    randomFactUrl := fmt.Sprintf("%s/randomfact/", WhiskeyApiUrl)
    // make the http request
    resp, err := http.Get(randomFactUrl)
    if err != nil {
        log.Fatalln(err)
    }

    // parse the response body
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatalln(err)
    }

    // decode the response json string
    apiResponse := RandomFactApiResponse{}
    err = json.Unmarshal(body, &apiResponse)
    if err != nil {
        log.Fatalln(err)
    }

    return string(apiResponse.Results[0].Text)
}


