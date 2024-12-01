package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/watevadafa/retrospec/internal/types"
	"strings"
)

// WebErrorHandler handles errors for web routes by rendering error templates
func WebErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	return c.Status(code).Render(fmt.Sprintf("errors/%d", code), fiber.Map{
		"Title": fmt.Sprintf("Error %d", code),
		"Error": err.Error(),
	})
}

// APIErrorHandler handles errors for API routes by returning JSON responses
func APIErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	if error, ok := err.(*fiber.Error); ok {
		code = error.Code
	}

	return c.Status(code).JSON(types.Error{
		Code:    code,
		Message: err.Error(),
	})
}

// ErrorHandler decides which error handler to use based on the path
func ErrorHandler(c *fiber.Ctx, err error) error {
	// Check if the request path starts with /api
	if strings.HasPrefix(c.Path(), "/api") {
		return APIErrorHandler(c, err)
	}
	return WebErrorHandler(c, err)
}
