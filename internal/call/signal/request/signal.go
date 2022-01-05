package request

type Signal struct {
	Type     string      `json:"type,omitempty"`
	PeerName string      `json:"peer_name,omitempty"`
	Payload  interface{} `json:"payload,omitempty"`
}
