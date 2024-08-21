package converter

import (
	"chatmoon/internal/entity"
	"chatmoon/internal/model"
)

func ChatHistoryToResponse(chatHistory *entity.ChatHistory) *model.ChatHistoryResponse {
	return &model.ChatHistoryResponse{
		ID:        chatHistory.ID,
		RoomID:    chatHistory.RoomID,
		UserID:    chatHistory.UserID,
		Message:   chatHistory.Message,
		CreatedAt: chatHistory.CreatedAt,
		UpdatedAt: chatHistory.UpdatedAt,
	}
}
