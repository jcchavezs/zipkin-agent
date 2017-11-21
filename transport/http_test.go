package transport

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jcchavezs/zipkin-agent"
	"github.com/magiconair/properties/assert"
)

const marshalledSpans = `[{"traceId":"352bff9a74ca9ad2","id":"6b221d5bc9e6496c","kind":"CLIENT"}]`

func TestGetParams(t *testing.T) {
	var payload []byte
	observableHandler := func(w http.ResponseWriter, r *http.Request) {
		payload, _ = ioutil.ReadAll(r.Body)
	}

	ts := httptest.NewServer(http.HandlerFunc(observableHandler))
	defer ts.Close()

	ht := NewHttpTransporter(ts.URL)
	ht.Send(&[]zipkinagent.Span{
		{TraceID: "352bff9a74ca9ad2", ID: "6b221d5bc9e6496c", Kind: "CLIENT"},
	})

	assert.Equal(t, string(payload), marshalledSpans)
}
