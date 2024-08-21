package socket

import (
	"chatmoon/internal/model"
	"chatmoon/internal/repository"
	"fmt"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const (
	MESSAGE_CHAT = "chat"
)

type ChatRoomSocket struct {
	DB             *gorm.DB
	Log            *logrus.Logger
	RoomRepository *repository.RoomRepository
	Connections    map[string][]*websocket.Conn
}

func NewChatRoomSocket(db *gorm.DB, log *logrus.Logger, RoomRepository *repository.RoomRepository) *ChatRoomSocket {
	return &ChatRoomSocket{
		DB:             db,
		Log:            log,
		RoomRepository: RoomRepository,
		Connections:    make(map[string][]*websocket.Conn, 0),
	}
}

func (c *ChatRoomSocket) ChitChatHandler(ctx *fiber.Ctx) error {
	if !websocket.IsWebSocketUpgrade(ctx) {
		return fiber.ErrUpgradeRequired
	}

	roomId := ctx.Params("room_id")
	count, err := c.RoomRepository.CountById(c.DB, roomId)
	if err != nil || count == 0 {
		return fiber.ErrNotFound
	}

	return ctx.Next()
}

func (c *ChatRoomSocket) ChitChatSocket(conn *websocket.Conn) {
	roomId := conn.Params("room_id")
	user := conn.Locals("auth").(*model.Auth)

	payload := &model.ChatSocketPayload{
		From:   user.ID,
		RoomID: roomId,
	}

	c.HandleOnline(conn, payload)

	conn.SetCloseHandler(func(code int, text string) error {
		c.HandleOffline(conn, payload)
		c.Log.Debug("Connections", c.Connections)
		return nil
	})

	for {
		if err := conn.ReadJSON(payload); err != nil {
			c.Log.Debug("Error read json : ", err)
			break
		}
		switch payload.Type {
		case MESSAGE_CHAT:
			c.HandleBroadcastMessage(payload)
		}
	}
}

func (c *ChatRoomSocket) ChitChatRecover(conn *websocket.Conn) {
	if err := recover(); err != nil {
		c.Log.WithError(err.(error)).Error("Error recover : ", err)
	}
}

func (c *ChatRoomSocket) HandleOnline(conn *websocket.Conn, payload *model.ChatSocketPayload) {
	c.Connections[payload.RoomID] = append(c.Connections[payload.RoomID], conn)
	c.HandleBroadcastMessage(&model.ChatSocketPayload{
		From:    "System",
		RoomID:  payload.RoomID,
		Message: fmt.Sprintf("%s online", payload.From),
	})
}

func (c *ChatRoomSocket) HandleOffline(currenConn *websocket.Conn, payload *model.ChatSocketPayload) {
	c.HandleBroadcastMessage(&model.ChatSocketPayload{
		From:    "System",
		RoomID:  payload.RoomID,
		Message: fmt.Sprintf("%s offline", payload.From),
	})

	for i, conn := range c.Connections[payload.RoomID] {
		if conn == currenConn {
			c.Connections[payload.RoomID] = append(c.Connections[payload.RoomID][:i], c.Connections[payload.RoomID][i+1:]...)
		}
	}

	if len(c.Connections[payload.RoomID]) == 0 {
		delete(c.Connections, payload.RoomID)
	}
}

func (c *ChatRoomSocket) HandleBroadcastMessage(payload *model.ChatSocketPayload) {
	for _, conn := range c.Connections[payload.RoomID] {
		if err := conn.WriteJSON(&model.ChatSocketResponse{
			From:    payload.From,
			Message: payload.Message,
		}); err != nil {
			c.Log.Debug("Error write json : ", err)
		}
	}
}
