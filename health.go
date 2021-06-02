package healthcheck

import (
	"github.com/gofiber/fiber/v2"
)

func New(config ...Config) fiber.Handler {
	// set default config
	cfg := configDefault(config...)

	return func(c *fiber.Ctx) error {

		value := c.Request().Header.Peek(cfg.HeaderName)
		if string(value) == cfg.HeaderValue {
			return c.Status(cfg.ResponseCode).SendString(cfg.ResponseText)
		}

		return c.Next()
	}
}
