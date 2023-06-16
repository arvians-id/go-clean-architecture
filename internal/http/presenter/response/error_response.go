package response

import "github.com/gofiber/fiber/v2"

var (
	NotFound = "record not found"
)

func ReturnErrorValidation(c *fiber.Ctx, code int, err []*ErrorResponse) error {
	return c.Status(code).JSON(ErrorValidationResponse{
		Code:   code,
		Status: "There are errors validation",
		Error:  err,
	})
}

func ReturnError(c *fiber.Ctx, code int, err error) error {
	return c.Status(code).JSON(ApiResponse{
		Code:   code,
		Status: err.Error(),
		Data:   nil,
	})
}
