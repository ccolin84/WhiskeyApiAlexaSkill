package main

type AlexaResponse struct {
	Version string  `json:"version"`
	Body    AlexaResBody `json:"response"`
}

type AlexaResBody struct {
	OutputSpeech     AlexaPayload `json:"outputSpeech,omitempty"`
	ShouldEndSession bool    `json:"shouldEndSession"`
}

type AlexaPayload struct {
	Type string `json:"type,omitempty"`
	Text string `json:"text,omitempty"`
}

func NewAlexaResponse(speech string) AlexaResponse {
	return AlexaResponse{
        Version: "1.0",
        Body: AlexaResBody{
            OutputSpeech: AlexaPayload{
                Type: "PlainText",
                Text: speech,
            },
            ShouldEndSession: true,
        },
    }
}
