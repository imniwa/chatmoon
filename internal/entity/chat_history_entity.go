package entity

import "time"

type ChatHistory struct {
	ID        string    `gorm:"column:id"`
	RoomID    string    `gorm:"foreignkey:rooms;references:room_id"`
	UserID    string    `gorm:"foreignkey:users;references:user_id"`
	Message   string    `gorm:"column:message"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime:true"`
	UpdateAt  time.Time `gorm:"column:updated_at;autoCreateTime:true;autoUpdateTime:true"`
}

func (ch *ChatHistory) TableName() string {
	return "chat_histories"
}
