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

type RoomUseCase struct {
	DB             *gorm.DB
	Log            *logrus.Logger
	Validate       *validator.Validate
	RoomRepository *repository.RoomRepository
}

func NewRoomUseCase(db *gorm.DB, log *logrus.Logger, validate *validator.Validate, roomRepository *repository.RoomRepository) *RoomUseCase {
	return &RoomUseCase{
		DB:             db,
		Log:            log,
		Validate:       validate,
		RoomRepository: roomRepository,
	}
}

func (c *RoomUseCase) Create(ctx context.Context, request *model.CreateRoomRequest) (*model.RoomResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	room := &entity.Room{
		ID:          uuid.New().String(),
		DisplayName: request.DisplayName,
		Description: request.Description,
		CreatedBy:   request.UserId,
	}

	if err := c.RoomRepository.Create(tx, room); err != nil {
		c.Log.WithError(err).Error("error create room")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error create room")
		return nil, fiber.ErrInternalServerError
	}

	return converter.RoomToResponse(room), nil
}

func (c *RoomUseCase) Search(ctx context.Context, request *model.SearchRoomRequest) ([]model.RoomResponse, int64, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, 0, fiber.ErrBadRequest
	}

	rooms, total, err := c.RoomRepository.Search(tx, request)

	if err != nil {
		c.Log.WithError(err).Error("error getting room")
		return nil, 0, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error getting room")
		return nil, 0, fiber.ErrInternalServerError
	}

	responses := make([]model.RoomResponse, len(rooms))

	for i, room := range rooms {
		responses[i] = *converter.RoomToResponse(&room)
	}

	return responses, total, nil
}
