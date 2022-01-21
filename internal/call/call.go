package call

import (
	"fmt"

	"github.com/mehditeymorian/webrtc-video-call/internal/call/signal"
	"github.com/mehditeymorian/webrtc-video-call/internal/call/signal/request"
	"github.com/mehditeymorian/webrtc-video-call/internal/call/signal/response"
	"github.com/mehditeymorian/webrtc-video-call/internal/model"
	"github.com/mehditeymorian/webrtc-video-call/internal/store"

	"github.com/gofiber/websocket/v2"
)

type RoomManager struct {
	Store  store.Store
	Room   *model.Room
	PeerID string
}

func NewPeerManager(store store.Store, room *model.Room) *RoomManager {
	return &RoomManager{Store: store, Room: room}
}

func (m *RoomManager) HandleRequest(req request.Signal, c *websocket.Conn) error {
	switch req.Type {
	case signal.JoinRequest:
		return m.handlePeerJoin(req, c)
	case signal.SdpOffer, signal.SdpAnswer:
		return m.handleSdp(req)
	case signal.IceCandidate:
		return m.handleIceCandidate(req)

	}

	return nil
}

func (m *RoomManager) handlePeerJoin(req request.Signal, c *websocket.Conn) error {
	peerID, err := m.Store.RoomCollection.AddPeer(m.Room.ID, req.PeerName, c)
	if err != nil {
		return fmt.Errorf("failed to add peer: %v", err)
	}

	m.PeerID = peerID

	signalResponse := response.NewSignalResponse(signal.JoinResponse, peerID)

	err = c.WriteJSON(signalResponse)
	if err != nil {
		return fmt.Errorf("failed to response join: %v", err)
	}

	return nil
}

func (m *RoomManager) handleSdp(req request.Signal) error {
	sdp, ok := req.Payload.(map[string]interface{})
	if !ok {
		return fmt.Errorf("failed to read sdp offer")
	}

	peer, err := m.Store.RoomCollection.GetPeer(m.Room.ID, sdp["destination_id"].(string))
	if err != nil {
		return fmt.Errorf("failed to send sdp to peer: %v", err)
	}

	err = peer.Connection.WriteJSON(req)
	if err != nil {
		return fmt.Errorf("failed to exchange sdp: %v", err)
	}

	return nil
}

func (m *RoomManager) handleIceCandidate(req request.Signal) error {
	iceCandidate, ok := req.Payload.(map[string]interface{})
	if !ok {
		return fmt.Errorf("failed to read sdp offer")
	}

	peer, err := m.Store.RoomCollection.GetPeer(m.Room.ID, iceCandidate["destination_id"].(string))
	if err != nil {
		return fmt.Errorf("failed to send sdp to peer: %v", err)
	}

	err = peer.Connection.WriteJSON(req)
	if err != nil {
		return fmt.Errorf("failed to exchange sdp: %v", err)
	}

	return nil
}
