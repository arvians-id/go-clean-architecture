package middleware

import (
	"fmt"
	"github.com/arvians-id/go-clean-architecture/cmd/config"
	"github.com/arvians-id/go-clean-architecture/internal/http/presenter/response"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/utils"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"
)

var SecretKey = []byte("secret")

func XApiKeyMiddleware(configuration config.Config) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		xApiKey := ctx.Get("X-API-KEY")
		if xApiKey == "" {
			ctx.Locals("middleware", "XApiKey Middleware")
			return fiber.NewError(fiber.StatusForbidden, "access denied: please provide a valid API key to access this page.")
		}

		if xApiKey != configuration.Get("X_API_KEY") {
			ctx.Locals("middleware", "XApiKey Middleware")
			return fiber.NewError(fiber.StatusForbidden, "invalid key: the provided API key is incorrect. please make sure to use a valid API key to access this resource.")
		}

		return ctx.Next()
	}
}

func NewJWTMiddleware() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: SecretKey,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return response.ReturnError(ctx, fiber.StatusUnauthorized, err)
		},
		SuccessHandler: func(ctx *fiber.Ctx) error {
			userContext := ctx.Locals("user").(*jwt.Token)
			userClaims := userContext.Claims.(jwt.MapClaims)

			ctx.Locals("id", userClaims["id"])
			ctx.Locals("email", userClaims["email"])
			return ctx.Next()
		},
	})
}

func NewCORSMiddleware() fiber.Handler {
	return cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, X-API-KEY",
		AllowMethods:     "POST, DELETE, PUT, PATCH, GET",
		AllowCredentials: true,
	})
}

func NewLoggerMiddleware(logFile *os.File) fiber.Handler {
	return logger.New(logger.Config{
		Format:     "[${time}] | ${status} | ${latency} | ${ip} | ${method} | ${path} | ${error}\n",
		Output:     logFile,
		TimeFormat: "02-Jan-2006 15:04:05",
		Done: func(c *fiber.Ctx, logString []byte) {
			fmt.Print(string(logString))
		},
	})
}

func NewCSRFMiddleware() fiber.Handler {
	return csrf.New(csrf.Config{
		KeyLookup:      "header:X-CSRF-Token",
		CookieName:     "csrf_token",
		CookieSameSite: "Lax",
		CookieHTTPOnly: true,
		Expiration:     15 * time.Minute,
		KeyGenerator:   utils.UUID,
	})
}
