package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/midoon/e-wallet-go-app-v1/dto"
	"github.com/midoon/e-wallet-go-app-v1/helper"
	"github.com/midoon/e-wallet-go-app-v1/internal/config"
)

func AuthMiddleware(cnf *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := strings.ReplaceAll(c.Get("Authorization"), "Bearer ", "")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, helper.ErrJwtValidation
			} else if method != jwt.SigningMethodHS256 {
				return nil, helper.ErrJwtValidation
			}

			return []byte(cnf.JWT.Key), nil
		})
		if err != nil {
			return c.Status(helper.HttpStatusErr(err)).JSON(dto.ErrorResponse{
				Status:  false,
				Message: err.Error(),
			})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(helper.HttpStatusErr(err)).JSON(dto.ErrorResponse{
				Status:  false,
				Message: "error  jwt parse to claim",
			})
		}

		c.Locals("x-user-id", claims["Id"])

		return c.Next()
	}
}
