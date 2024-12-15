package prober

import "github.com/gofiber/fiber/v2"

func LaunchProbes(cfg Config) error {
	app := fiber.New(
		fiber.Config{
			DisableStartupMessage: true,
		},
	)

	app.Get(cfg.ReadinessPath, func(c *fiber.Ctx) error {
		return c.SendString("Ready")
	})

	app.Get(cfg.LivenessPath, func(c *fiber.Ctx) error {
		return c.SendString("Alive")
	})

	errChan := make(chan error, 1)
	go func() {
		if err := app.Listen(cfg.Address); err != nil {
			errChan <- err
		} else {
			errChan <- nil
		}
	}()

	return <-errChan
}
