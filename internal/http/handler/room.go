package handler

import (
	"log"
	"net/http"

	"hermes/internal/call"
	"hermes/internal/call/signal"
	"hermes/internal/call/signal/request"
	"hermes/internal/call/signal/response"
	"hermes/internal/store"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

type Room struct {
	Store store.Store
}

func (r Room) Register(app *fiber.App) {
	app.Post("/room", r.create)
	app.Get("/ws/room/:room_id", websocket.New(r.join))

	app.Static("/", "ui/index.html", fiber.Static{
		Browse: true,
	})
	app.Static("/room", "ui/room.html", fiber.Static{
		Browse: true,
	})

	app.Static("/not-found", "ui/404.html", fiber.Static{
		Browse: true,
	})
}

func (r Room) create(ctx *fiber.Ctx) error {

	roomID := r.Store.RoomCollection.Create()

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"room_id": roomID,
	})
}

func (r Room) join(c *websocket.Conn) {

	roomID := c.Params("room_id")

	room, err := r.Store.RoomCollection.Get(roomID)
	if err != nil {
		log.Printf("failed to join the room: %v", err)
		notFound := response.NewSignalResponse(signal.RoomNotFound, "")
		c.Conn.WriteJSON(notFound)
		c.Conn.Close()

		return
	}

	peerManager := call.NewPeerManager(r.Store, room)

	for {
		var req request.Signal

		err := c.ReadJSON(&req)
		if err != nil {
			log.Printf("failed to read json %v", err)
			room.OnLeave(peerManager.PeerID)
			err := r.Store.RoomCollection.DelPeer(roomID, peerManager.PeerID)
			log.Printf("failed to remove peer: %v", err)
			break
		}

		log.Printf("request type: %s peerName: %s", req.Type, req.PeerName)

		err = peerManager.HandleRequest(req, c)
		if err != nil {
			log.Printf("failed to handle request: %v", err)
		}
	}

}
