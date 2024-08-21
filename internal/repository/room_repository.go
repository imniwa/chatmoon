package repository

import (
	"chatmoon/internal/entity"

	"github.com/sirupsen/logrus"
)

type RoomRepository struct {
	Repository[entity.Room]
	Log *logrus.Logger
}

func NewRoomRepository(log *logrus.Logger) *RoomRepository {
	return &RoomRepository{
		Log: log,
	}
}
