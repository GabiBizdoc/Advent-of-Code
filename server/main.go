package server

import (
	"aoc/server/config"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"log"
	"os"
	"os/signal"
)

func loadRoutes(app *fiber.App) {
	//db := db.CreateConnection(env.Config.DBConnectionString)
	//api := app.Group("/api")
	//hoarder.ApplyRoutes(api, db)
}

func Start() {
	env.LoadConfig()
	app := fiberApp()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		_ = <-c
		fmt.Println("Gracefully shutting down...")
		_ = app.Shutdown()
	}()

	log.Fatal(app.Listen(env.Config.AppHost + ":4001"))
}

// fiberApp calls loadRoutes internally
func fiberApp() *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}
			return ctx.Status(code).SendString(err.Error())
		},
	})

	app.Use(helmet.New())
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format: "${time} [${ip}]:${port} ${pid} ${locals:requestid} ${status} - ${latency} ${method} ${path}\n",
	}))

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("healthy!")
	})

	loadRoutes(app)
	return app
}
