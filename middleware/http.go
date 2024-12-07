package middleware

import (
	"time"

	"github.com/Markuysa/pkg/logger"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func AccessLogMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()
		duration := time.Since(start)

		logger.Logger.Info("HTTP request",
			zap.String("method", c.Method()),
			zap.String("url", c.OriginalURL()),
			zap.Duration("duration", duration),
			zap.Int("status", c.Response().StatusCode()),
		)

		return err
	}
}
