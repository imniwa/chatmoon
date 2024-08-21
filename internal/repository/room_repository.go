package repository

import (
	"chatmoon/internal/entity"
	"chatmoon/internal/model"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
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

func (r *RoomRepository) Search(db *gorm.DB, request *model.SearchRoomRequest) ([]entity.Room, int64, error) {
	var rooms []entity.Room
	if err := db.Scopes(r.FilterRoom(request)).Offset((request.Page - 1) * request.Size).Limit(request.Size).Find(&rooms).Error; err != nil {
		return nil, 0, err
	}

	var total int64 = 0
	if err := db.Model(&entity.Room{}).Scopes(r.FilterRoom(request)).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return rooms, total, nil
}

func (r *RoomRepository) FilterRoom(request *model.SearchRoomRequest) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {

		if displayName := request.DisplayName; displayName != "" {
			tx = tx.Where("display_name LIKE ?", "%"+displayName+"%")
		}

		if description := request.Description; description != "" {
			tx = tx.Where("description LIKE ?", "%"+description+"%")
		}

		if id := request.ID; id != "" {
			tx = tx.Where("id = ?", id)
		}

		if userId := request.UserId; userId != "" {
			tx = tx.Where("created_by = ?", userId)
		}

		return tx
	}
}
