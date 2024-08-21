package model

import "time"

// Response
type ChatHistoryResponse struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	RoomID    string    `json:"room_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Request
type InsertChatHistoryRequest struct {
	UserID  string `json:"-" validate:"required"`
	RoomID  string `json:"-" validate:"required"`
	Message string `json:"message" validate:"required"`
}

type SearchChatHistoryRequest struct {
	ID      string `json:"id" validate:"max=32"`
	UserID  string `json:"user_id" validate:"max=32"`
	RoomID  string `json:"room_id" validate:"max=32"`
	Message string `json:"message" validate:"max=128"`
	Page    int    `json:"page" validate:"min=1"`
	Size    int    `json:"size" validate:"min=1,max=100"`
}
