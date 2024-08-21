package http

import (
	"chatmoon/internal/delivery/http/middleware"
	"chatmoon/internal/model"
	"chatmoon/internal/usecase"
	"math"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type RoomController struct {
	Log     *logrus.Logger
	UseCase *usecase.RoomUseCase
}

func NewRoomController(log *logrus.Logger, useCase *usecase.RoomUseCase) *RoomController {
	return &RoomController{
		Log:     log,
		UseCase: useCase,
	}
}

func (c *RoomController) Create(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := new(model.CreateRoomRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.Warnf("Failed to parse request: %v", err)
		return fiber.ErrBadRequest
	}
	request.UserId = auth.ID

	response, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to create room: %v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[*model.RoomResponse]{Data: response})
}

func (c *RoomController) List(ctx *fiber.Ctx) error {
	request := &model.SearchRoomRequest{
		ID:          ctx.Query("id", ""),
		UserId:      ctx.Query("user_id", ""),
		DisplayName: ctx.Query("display_name", ""),
		Description: ctx.Query("description", ""),
		Page:        ctx.QueryInt("page", 1),
		Size:        ctx.QueryInt("size", 10),
	}

	responses, total, err := c.UseCase.Search(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("Failed to search room")
		return err
	}

	paging := &model.PageMetaData{
		Page:      request.Page,
		Size:      request.Size,
		TotalItem: total,
		TotalPage: int64(math.Ceil(float64(total) / float64(request.Size))),
	}

	return ctx.JSON(model.WebResponse[[]model.RoomResponse]{
		Data:   responses,
		Paging: paging,
	})
}
