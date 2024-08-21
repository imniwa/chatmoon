package entity

import "time"

type Room struct {
	ID          string    `gorm:"column:id"`
	DisplayName string    `gorm:"column:display_name"`
	Description string    `gorm:"column:description"`
	CreatedBy   string    `gorm:"column:created_by"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime:true"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoCreateTime:true;autoUpdateTime:true"`
}

func (r *Room) TableName() string {
	return "rooms"
}
