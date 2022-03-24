package response

import (
	"github.com/gofiber/fiber/v2"
)

// struct to store the response or map the response to send as output
type res struct {
	Status     interface{} `json:"status"`
	Message    interface{} `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	Error_obj  interface{} `json:"error_obj,omitempty"`
	Pagination interface{} `json:"pagination,omitempty"`
}

// To send the response
func Response(c *fiber.Ctx) error {

	var result res

	result.Status = c.Locals("status")
	result.Message = c.Locals("message")
	result.Data = c.Locals("data")
	result.Error_obj = c.Locals("error_obj")
	result.Pagination = c.Locals("pagination")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": result,
	})
}
