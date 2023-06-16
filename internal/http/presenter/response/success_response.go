package response

import "github.com/gofiber/fiber/v2"

func ReturnSuccess(c *fiber.Ctx, code int, status string, data interface{}) error {
	return c.Status(code).JSON(ApiResponse{
		Code:   code,
		Status: status,
		Data:   data,
	})
}
