package store

import (
	"github.com/mehditeymorian/webrtc-video-call/internal/model"
	"github.com/mehditeymorian/webrtc-video-call/internal/store/collection"
)

type Store struct {
	RoomCollection collection.Room
}

func Create() Store {
	return Store{
		RoomCollection: collection.Room{
			Rooms: make(map[string]*model.Room),
		},
	}
}
