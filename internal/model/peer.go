package model

import (
	"sync"

	"github.com/gofiber/websocket/v2"
)

type Peer struct {
	LocK       *sync.Mutex
	Connection *websocket.Conn
	ID         string
	Name       string
}

func NewPeer(connection *websocket.Conn, ID string, name string) *Peer {
	return &Peer{
		LocK:       &sync.Mutex{},
		Connection: connection,
		ID:         ID,
		Name:       name,
	}
}
