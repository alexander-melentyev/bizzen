package respfmt

// Fmt -.
type Fmt struct {
	ServiceName string      `json:"serviceName,omitempty"`
	UserMsg     string      `json:"userMessage,omitempty"`
	DevMsg      string      `json:"developerMessage,omitempty"`
	Data        interface{} `json:"data,omitempty"`
	TotalCount  int64       `json:"totalCount,omitempty"`
}
