package event

import (
	"log"

	"hermes/internal/call/signal"
	"hermes/internal/call/signal/response"
	"hermes/internal/model"

	"github.com/gofiber/fiber/v2"
)

func RoomOnJoin(room *model.Room) func(peerID string) {
	return func(peerID string) {
		peerName := room.Peers[peerID].Name
		log.Printf("onJoin: %s %s", peerID, peerName)

		resp := response.NewSignalResponse(signal.NewPeer, fiber.Map{
			"name": peerName,
			"id":   peerID,
		})

		broadcastSignal(room, peerID, resp)
	}
}

func RoomOnLeave(room *model.Room) func(peerID string) {
	return func(peerID string) {
		resp := response.NewSignalResponse(signal.PeerLeave, peerID)
		broadcastSignal(room, peerID, resp)
	}
}

func RoomOnClose(room *model.Room) func() {
	return func() {
		resp := response.NewSignalResponse(signal.RoomClose, nil)
		broadcastSignal(room, "", resp)
	}
}

func broadcastSignal(room *model.Room, exceptID string, signal *response.Signal) {
	for id, peer := range room.Peers {
		if id == exceptID {
			continue
		}

		peer.LocK.Lock()
		if err := peer.Connection.WriteJSON(signal); err != nil {
			log.Printf("failed to send signal of type %s to %s: %v", signal.Type, id, err)
		}
		peer.LocK.Unlock()
	}
}
