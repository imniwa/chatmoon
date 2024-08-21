package config

import (
	"chatmoon/internal/delivery/http"
	"chatmoon/internal/delivery/http/middleware"
	"chatmoon/internal/delivery/http/route"
	"chatmoon/internal/delivery/http/socket"
	"chatmoon/internal/repository"
	"chatmoon/internal/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB       *gorm.DB
	App      *fiber.App
	Log      *logrus.Logger
	Validate *validator.Validate
	Config   *viper.Viper
}

func Bootstrap(config *BootstrapConfig) {
	// Setup Repository
	userRepository := repository.NewUserRepository(config.Log)
	roomRepository := repository.NewRoomRepository(config.Log)
	chatHistoryRepository := repository.NewChatHistoryRepository(config.Log)

	// Setup Usecase
	userUsecase := usecase.NewUserUseCase(config.DB, config.Log, config.Validate, userRepository)
	roomUseCase := usecase.NewRoomUseCase(config.DB, config.Log, config.Validate, roomRepository)
	chatHistoryUseCase := usecase.NewChatHistoryUseCase(config.DB, config.Log, config.Validate, chatHistoryRepository)

	// Setup Controller
	userController := http.NewUserController(config.Log, userUsecase)
	roomController := http.NewRoomController(config.Log, roomUseCase)
	chatHistoryController := http.NewChatHistoryController(config.Log, chatHistoryUseCase)

	// Setup Socket
	chatRoomSocket := socket.NewChatRoomSocket(config.Log, chatHistoryUseCase, roomUseCase)

	authMiddleware := middleware.NewAuth(userUsecase)

	routeConfig := &route.RouteConfig{
		App:                   config.App,
		UserController:        userController,
		RoomController:        roomController,
		ChatHistoryController: chatHistoryController,
		ChatRoomSocket:        chatRoomSocket,
		AuthMiddleware:        authMiddleware,
	}
	routeConfig.Setup()
}
