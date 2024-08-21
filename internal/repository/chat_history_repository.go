package repository

import (
	"chatmoon/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ChatHistoryRepository struct {
	Repository[entity.ChatHistory]
	Log *logrus.Logger
}

func NewChatHistoryRepository(log *logrus.Logger) *ChatHistoryRepository {
	return &ChatHistoryRepository{
		Log: log,
	}
}

func (chr *ChatHistoryRepository) FindByRoomID(db *gorm.DB, chatHistory *entity.ChatHistory, roomID string) error {
	return db.Where("room_id = ?", roomID).First(chatHistory).Error
}
