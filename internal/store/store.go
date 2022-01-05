package store

import (
	"hermes/internal/model"
	"hermes/internal/store/collection"
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
