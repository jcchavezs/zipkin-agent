package zipkinagent

import (
	"encoding/json"
	"log"
)

type Transporter interface {
	Send(*[]Span)
}

type NoopTransport struct{}

func (t *NoopTransport) Send(_ *[]Span) {}

type loggerTransport struct{}

func NewLoggerTransporter() Transporter {
	return &loggerTransport{}
}

func (t *loggerTransport) Send(spans *[]Span) {
	if b, err := json.MarshalIndent(spans, "", "  "); err == nil {
		log.Printf("%s\n", string(b))
	}
}
