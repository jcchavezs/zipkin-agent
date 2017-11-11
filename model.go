package zipkinagent

type Endpoint struct {
	ServiceName string `json:"serviceName,omitempty"`
	IPv4        string `json:"ipv4,omitempty"`
	Port        int    `json:"port,omitempty"`
}

type Log struct {
	Timestamp int64  `json:"timestamp,omitempty"`
	Value     string `json:"value,omitempty"`
}

type Span struct {
	Kind            string            `json:"kind,omitempty"`
	TraceID         string            `json:"traceId,omitempty"`
	ParentID        string            `json:"parentId,omitempty"`
	ID              string            `json:"id,omitempty"`
	Name            string            `json:"name,omitempty"`
	Timestamp       int64             `json:"timestamp,omitempty"`
	TimestampMicros int64             `json:"timestampMicros,omitempty"`
	Duration        int64             `json:"duration,omitempty"`
	DurationMicros  int64             `json:"durationMicros,omitempty"`
	LocalEndpoint   *Endpoint         `json:"localEndpoint,omitempty"`
	RemoteEndpoint  *Endpoint         `json:"remoteEndpoint,omitempty"`
	Tags            map[string]string `json:"tags,omitempty"`
	Logs            []Log             `json:"annotations,omitempty"`
}
