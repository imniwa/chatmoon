package repository

import (
	"chatmoon/internal/entity"
	"chatmoon/internal/model"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func (chr *ChatHistoryRepository) Search(db *gorm.DB, request *model.SearchChatHistoryRequest) ([]entity.ChatHistory, int64, error) {
	var chats []entity.ChatHistory
	if err := db.Scopes(chr.FilterChat(request)).Offset((request.Page - 1) * request.Size).Limit(request.Size).Order(clause.OrderByColumn{
		Column: clause.Column{Name: "created_at"},
		Desc:   true,
	}).Find(&chats).Error; err != nil {
		return nil, 0, err
	}

	var total int64 = 0
	if err := db.Model(&entity.ChatHistory{}).Scopes(chr.FilterChat(request)).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return chats, total, nil
}

func (chr *ChatHistoryRepository) FilterChat(request *model.SearchChatHistoryRequest) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {

		if id := request.ID; id != "" {
			tx = tx.Where("id = ?", id)
		}

		if userId := request.UserID; userId != "" {
			tx = tx.Where("user_id = ?", userId)
		}

		if roomId := request.RoomID; roomId != "" {
			tx = tx.Where("room_id = ?", roomId)
		}

		if message := request.Message; message != "" {
			tx = tx.Where("message LIKE ?", "%"+message+"%")
		}

		return tx
	}
}
