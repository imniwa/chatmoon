package middleware

import (
	"chatmoon/internal/model"
	"chatmoon/internal/usecase"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func NewAuth(userUserCase *usecase.UserUseCase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		request := &model.VerifyUserRequest{Token: ctx.Get("Authorization", "NOT_FOUND")}
		token := strings.Fields(request.Token)
		userUserCase.Log.Debugf("Authorization : %s", token)

		if len(token) != 2 || token[0] != "Bearer" {
			userUserCase.Log.Warnf("Failed find user by token : %+v", token)
			return fiber.ErrUnauthorized
		}

		auth, err := userUserCase.Verify(ctx.UserContext(), &model.VerifyUserRequest{Token: token[1]})
		if err != nil {
			userUserCase.Log.Warnf("Failed find user by token : %+v", err)
			return fiber.ErrUnauthorized
		}

		ctx.Locals("auth", auth)
		return ctx.Next()
	}
}

func GetUser(ctx *fiber.Ctx) *model.Auth {
	return ctx.Locals("auth").(*model.Auth)
}
