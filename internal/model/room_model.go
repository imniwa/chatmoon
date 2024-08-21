package model

import "time"

// Response
type RoomResponse struct {
	ID          string    `json:"id,omitempty"`
	DisplayName string    `json:"display_name,omitempty"`
	Description string    `json:"description,omitempty"`
	CreatedBy   string    `json:"created_by,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

// Request
type CreateRoomRequest struct {
	UserId      string `json:"-" validate:"required,max=32"`
	DisplayName string `json:"display_name" validate:"required,max=64"`
	Description string `json:"description" validate:"max=128"`
}

type GetRoomRequest struct {
	ID string `json:"id" validate:"required,max=32"`
}

type SearchRoomRequest struct {
	UserId      string `json:"-" validate:"max=32"`
	ID          string `json:"id" validate:"max=32"`
	DisplayName string `json:"display_name" validate:"max=64"`
	Description string `json:"description" validate:"max=128"`
	Page        int    `json:"page" validate:"min=1"`
	Size        int    `json:"size" validate:"min=1,max=100"`
}

type UpdateRoomRequest struct {
	UserId      string `json:"-" validate:"required,max=32"`
	ID          string `json:"id" validate:"required,max=32"`
	DisplayName string `json:"display_name" validate:"required,max=64"`
	Description string `json:"description" validate:"required,max=128"`
}

type DeleteRoomRequest struct {
	ID string `json:"id" validate:"required,max=32"`
}
