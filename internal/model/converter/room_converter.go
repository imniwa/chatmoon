package converter

import (
	"chatmoon/internal/entity"
	"chatmoon/internal/model"
)

func RoomToResponse(room *entity.Room) *model.RoomResponse {
	return &model.RoomResponse{
		ID:          room.ID,
		DisplayName: room.DisplayName,
		Description: room.Description,
		CreatedAt:   room.CreatedAt,
		UpdatedAt:   room.UpdatedAt,
	}
}
