package http

import (
	"chatmoon/internal/delivery/http/middleware"
	"chatmoon/internal/model"
	"chatmoon/internal/usecase"
	"math"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type ChatHistoryController struct {
	Log     *logrus.Logger
	UseCase *usecase.ChatHistoryUseCase
}

func NewChatHistoryController(log *logrus.Logger, useCase *usecase.ChatHistoryUseCase) *ChatHistoryController {
	return &ChatHistoryController{
		Log:     log,
		UseCase: useCase,
	}
}

func (c *ChatHistoryController) FindByRoomID(ctx *fiber.Ctx) error {
	request := model.SearchChatHistoryRequest{
		RoomID: ctx.Params("room_id"),
		Page:   ctx.QueryInt("page", 1),
		Size:   ctx.QueryInt("size", 10),
	}

	responses, total, err := c.UseCase.FindByRoomID(ctx.UserContext(), &request)
	if err != nil {
		c.Log.WithError(err).Error("Failed to find chat history by room id")
		return err
	}

	paging := &model.PageMetaData{
		Page:      request.Page,
		Size:      request.Size,
		TotalItem: total,
		TotalPage: int64(math.Ceil(float64(total) / float64(request.Size))),
	}

	return ctx.JSON(model.WebResponse[[]model.ChatHistoryResponse]{
		Data:   responses,
		Paging: paging,
	})
}

func (c *ChatHistoryController) Insert(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := new(model.InsertChatHistoryRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.Warnf("Failed to parse request: %v", err)
		return fiber.ErrBadRequest
	}
	request.UserID = auth.ID
	request.RoomID = ctx.Params("room_id")

	response, err := c.UseCase.Insert(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to insert chat history: %v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[*model.ChatHistoryResponse]{Data: response})
}
