package utils

import (
	"github.com/gofiber/fiber/v2"
	"strings"
)

func ParseRequestBody(c *fiber.Ctx, out interface{}) error {
	contentType := string(c.Request().Header.ContentType())
	if strings.Contains(contentType, "application/x-www-form-urlencoded") || strings.Contains(contentType, "application/json") {
		return c.BodyParser(out)
	}
	return fiber.NewError(fiber.StatusUnsupportedMediaType, "Unsupported content type")
}
