package response

type Signal struct {
	Type    string      `json:"type,omitempty"`
	Payload interface{} `json:"payload,omitempty"`
}

func NewSignalResponse(signalType string, payload interface{}) *Signal {
	return &Signal{Type: signalType, Payload: payload}
}
