package config

import (
	"chatmoon/internal/delivery/http"
	"chatmoon/internal/delivery/http/middleware"
	"chatmoon/internal/delivery/http/route"
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

	// Setup Usecase
	userUsecase := usecase.NewUserUseCase(config.DB, config.Log, config.Validate, userRepository)

	// Setup Controller
	userController := http.NewUserController(config.Log, userUsecase)

	authMiddleware := middleware.NewAuth(userUsecase)

	routeConfig := &route.RouteConfig{
		App:            config.App,
		UserController: userController,
		AuthMiddleware: authMiddleware,
	}
	routeConfig.Setup()
}
