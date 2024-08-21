package usecase

import (
	"chatmoon/internal/entity"
	"chatmoon/internal/model"
	"chatmoon/internal/model/converter"
	"chatmoon/internal/repository"
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ChatHistoryUseCase struct {
	DB                    *gorm.DB
	Log                   *logrus.Logger
	Validate              *validator.Validate
	ChatHistoryRepository *repository.ChatHistoryRepository
}

func NewChatHistoryUseCase(db *gorm.DB, log *logrus.Logger, validate *validator.Validate, chatHistoryRepository *repository.ChatHistoryRepository) *ChatHistoryUseCase {
	return &ChatHistoryUseCase{
		DB:                    db,
		Log:                   log,
		Validate:              validate,
		ChatHistoryRepository: chatHistoryRepository,
	}
}

func (c *ChatHistoryUseCase) FindByRoomID(ctx context.Context, request *model.SearchChatHistoryRequest) ([]model.ChatHistoryResponse, int64, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	chats, total, err := c.ChatHistoryRepository.Search(tx, request)

	if err != nil {
		c.Log.WithError(err).Error("error search chat history")
		return nil, 0, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error search chat history")
		return nil, 0, fiber.ErrInternalServerError
	}

	responses := make([]model.ChatHistoryResponse, len(chats))
	for i, chat := range chats {
		responses[i] = *converter.ChatHistoryToResponse(&chat)
	}

	return responses, total, nil
}

func (c *ChatHistoryUseCase) Insert(ctx context.Context, request *model.InsertChatHistoryRequest) (*model.ChatHistoryResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	chatHistory := &entity.ChatHistory{
		ID:      uuid.New().String(),
		UserID:  request.UserID,
		RoomID:  request.RoomID,
		Message: request.Message,
	}

	if err := c.ChatHistoryRepository.Create(tx, chatHistory); err != nil {
		c.Log.WithError(err).Error("error insert chat history")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error insert chat history")
		return nil, fiber.ErrInternalServerError
	}

	return converter.ChatHistoryToResponse(chatHistory), nil
}
