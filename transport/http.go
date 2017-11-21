package transport

import (
	"encoding/json"
	"net/http"
	"bytes"
	"log"
	"github.com/jcchavezs/zipkin-agent"
)

type httpTransport struct {
	httpClient *http.Client
	url string
}

func (t *httpTransport) Send(spans *[]zipkinagent.Span) {
	sp, err := json.Marshal(spans)
	if err != nil {
		log.Printf("failed to marshall spans: %s\n", err.Error())
		return
	}

	res, err := t.httpClient.Post(t.url, "application/json", bytes.NewReader(sp))
	if err != nil {
		log.Printf("failed to send spans to the server: %s\n", err.Error())
		return
	}

	if res.StatusCode < 200 || res.StatusCode > 299 {
		log.Printf("invalid status code %d from server\n", res.StatusCode)
		return
	}
}

func NewHttpTransporter(httpServerURL string) zipkinagent.Transporter {
	return &httpTransport{http.DefaultClient, httpServerURL}
}
