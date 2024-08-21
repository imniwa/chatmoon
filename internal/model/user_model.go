package model

import "time"

// Response
type UserResponse struct {
	ID          string    `json:"id,omitempty"`
	DisplayName string    `json:"display_name,omitempty"`
	Token       string    `json:"token,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

// Request
type VerifyUserRequest struct {
	Token string `json:"token" validate:"required,max=128"`
}

type RegisterUserRequest struct {
	ID          string `json:"id" validate:"required,max=32"`
	DisplayName string `json:"display_name" validate:"required,max=64"`
	Password    string `json:"password" validate:"required,max=64"`
}

type UpdateUserRequest struct {
	ID          string `json:"id" validate:"required,max=32"`
	DisplayName string `json:"display_name" validate:"required,max=64"`
	Password    string `json:"password" validate:"required,max=128"`
}

type LoginUserRequest struct {
	ID       string `json:"id" validate:"required,max=32"`
	Password string `json:"password" validate:"required,max=128"`
}

type LogoutUserRequest struct {
	ID string `json:"id" validate:"required,max=32"`
}

type GetUserRequest struct {
	ID string `json:"id" validate:"required,max=32"`
}
