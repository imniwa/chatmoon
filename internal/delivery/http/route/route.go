package route

import (
	"chatmoon/internal/delivery/http"
	"chatmoon/internal/delivery/http/socket"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App                   *fiber.App
	UserController        *http.UserController
	RoomController        *http.RoomController
	ChatHistoryController *http.ChatHistoryController
	ChatRoomSocket        *socket.ChatRoomSocket
	AuthMiddleware        fiber.Handler
}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoute()
	c.SetupAuthRoute()
	c.SetupConfigSocket()
}

func (c *RouteConfig) SetupGuestRoute() {
	c.App.Post("/api/users", c.UserController.Register)
	c.App.Post("/api/users/_login", c.UserController.Login)
}

func (c *RouteConfig) SetupAuthRoute() {
	c.App.Use(c.AuthMiddleware)
	c.App.Delete("/api/users", c.UserController.Logout)
	c.App.Patch("/api/users/_current", c.UserController.Update)
	c.App.Get("/api/users/_current", c.UserController.Current)

	c.App.Post("/api/rooms", c.RoomController.Create)
	c.App.Get("/api/rooms", c.RoomController.List)

	c.App.Post("/api/rooms/:room_id/chats", c.ChatHistoryController.Insert)
	c.App.Get("/api/rooms/:room_id/chats", c.ChatHistoryController.FindByRoomID)
}

func (c *RouteConfig) SetupConfigSocket() {
	c.App.Get("/chatmoon/:room_id", c.ChatRoomSocket.ChitChatHandler, websocket.New(c.ChatRoomSocket.ChitChatSocket, websocket.Config{
		RecoverHandler: c.ChatRoomSocket.ChitChatRecover,
	}))
}
