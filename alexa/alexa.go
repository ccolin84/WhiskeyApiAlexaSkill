package alexa

type Response struct {
	Version string  `json:"version"`
	Body    ResBody `json:"response"`
}

type ResBody struct {
	OutputSpeech     Payload `json:"outputSpeech,omitempty"`
	ShouldEndSession bool    `json:"shouldEndSession"`
}

type Payload struct {
	Type string `json:"type,omitempty"`
	Text string `json:"text,omitempty"`
}

func NewResponse(speech string) Response {
	return Response{
        Version: "1.0",
        Body: ResBody{
            OutputSpeech: Payload{
                Type: "PlainText",
                Text: speech,
            },
            ShouldEndSession: true,
        },
    }
}
