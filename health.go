package healthcheck

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func New(config ...Config) fiber.Handler {
	// set default config
	cfg := configDefault(config...)

	return func(c *fiber.Ctx) error {

		header := make(http.Header)

		c.Request().Header.VisitAll(func(k, v []byte) {
			header.Set(string(k), string(v))
		})

		if header.Get(cfg.HeaderName) == cfg.HeaderValue {
			return c.Status(cfg.ResponseCode).SendString(cfg.ResponseText)
		}

		return c.Next()
	}
}
