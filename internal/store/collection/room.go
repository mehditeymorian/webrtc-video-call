package collection

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"hermes/internal/call/event"
	"hermes/internal/model"

	"github.com/gofiber/websocket/v2"
)

const (
	RoomIDLength = 8
	PeerIDLength = 16
)

type Room struct {
	Rooms map[string]*model.Room
}

// Create creates a room.
func (r *Room) Create() string {
	roomID := generateRandomHash(RoomIDLength)
	room := model.NewRoom(roomID)
	r.Rooms[roomID] = room

	room.OnJoin = event.RoomOnJoin(room)
	room.OnLeave = event.RoomOnLeave(room)
	room.OnClose = event.RoomOnClose(room)

	return roomID
}

// Get retrieve a room.
func (r *Room) Get(roomID string) (*model.Room, error) {
	room := r.Rooms[roomID]

	if room == nil {
		return nil, errors.New("room does not exists")
	}

	return room, nil
}

// Del delete a room and return its participants.
func (r *Room) Del(roomID string) error {
	room, err := r.Get(roomID)
	if err != nil {
		return fmt.Errorf("failed to find the room: %w", err)
	}

	for _, peer := range room.Peers {
		peer.Connection.Close()
	}

	room.OnClose()

	delete(r.Rooms, roomID)

	return nil
}

// AddPeer add a participant to a particular room.
func (r *Room) AddPeer(roomID, peerName string, conn *websocket.Conn) (string, error) {
	room, err := r.Get(roomID)
	if err != nil {
		return "", fmt.Errorf("failed to find the room: %w", err)
	}

	peerID := generateRandomHash(PeerIDLength)
	peer := model.NewPeer(conn, peerID, peerName)
	room.Peers[peerID] = peer

	room.OnJoin(peerID)

	return peerID, nil
}

func (r *Room) DelPeer(roomID, peerID string) error {
	room, err := r.Get(roomID)
	if err != nil {
		return fmt.Errorf("failed to find the room: %w", err)
	}

	delete(room.Peers, peerID)

	room.OnLeave(peerID)

	return nil
}

func (r Room) GetPeer(roomID, peerID string) (*model.Peer, error) {
	room, err := r.Get(roomID)
	if err != nil {
		return nil, fmt.Errorf("failed to find the room: %w", err)
	}

	peer := room.Peers[peerID]
	if peer == nil {
		return nil, fmt.Errorf("peer does not exists")
	}

	return peer, nil
}

// generateRandomHash produce random hashes with length of HashLength.
func generateRandomHash(length int) string {
	rand.Seed(time.Now().UnixNano())
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	b := make([]rune, length)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}
