package model

type Room struct {
	ID    string
	Peers map[string]*Peer

	OnJoin  func(peerID string)
	OnLeave func(peerID string)
	OnClose func()
}

func NewRoom(ID string) *Room {
	return &Room{
		ID:    ID,
		Peers: make(map[string]*Peer, 0),
	}
}
