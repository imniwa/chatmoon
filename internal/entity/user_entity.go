package entity

import "time"

type User struct {
	ID          string    `gorm:"column:id"`
	DisplayName string    `gorm:"column:display_name"`
	Password    string    `gorm:"column:password"`
	Token       string    `gorm:"column:token"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime:true"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoCreateTime:true;autoUpdateTime:true"`
}

func (u *User) TableName() string {
	return "users"
}
