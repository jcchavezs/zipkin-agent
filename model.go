package zipkinagent

type Endpoint struct {
	ServiceName string `json:"serviceName,omitempty"`
	IPv4        string `json:"ipv4,omitempty"`
	IPv6        string `json:"ipv6,omitempty"`
	Port        int    `json:"port,omitempty"`
}

type Annotation struct {
	Timestamp int64  `json:"timestamp,omitempty"`
	Value     string `json:"value,omitempty"`
}

type Span struct {
	TraceID        string            `json:"traceId"`
	Name           string            `json:"name,omitempty"`
	ParentID       string            `json:"parentId,omitempty"`
	ID             string            `json:"id"`
	Kind           string            `json:"kind,omitempty"`
	Timestamp      int64             `json:"timestamp,omitempty"`
	Duration       int64             `json:"duration,omitempty"`
	Debug          bool              `json:"debug"`
	Shared         bool              `json:"shared"`
	LocalEndpoint  *Endpoint         `json:"localEndpoint,omitempty"`
	RemoteEndpoint *Endpoint         `json:"remoteEndpoint,omitempty"`
	Annotations    []Annotation      `json:"annotations,omitempty"`
	Tags           map[string]string `json:"tags,omitempty"`
}
