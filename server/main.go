package server

import (
	"aoc/server/api"
	env "aoc/server/config"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/template/html/v2"
	"html/template"
	"log"
	"os"
	"os/signal"
	"time"
)

func Start() {
	env.LoadConfig()
	app := fiberApp()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		_ = <-c
		fmt.Println("Gracefully shutting down...")
		err := app.Shutdown()
		if err != nil {
			fmt.Println(err.Error())
		}
	}()

	log.Fatal(app.Listen(env.Config.AppHost + ":4001"))
}

// fiberApp calls loadRoutes internally
func fiberApp() *fiber.App {

	engine := html.New("./server/views/web", ".html")
	//engine.Reload(true)
	engine.AddFunc(
		"unsafe", func(s string) template.HTML {
			return template.HTML(s)
		},
	)

	app := fiber.New(fiber.Config{
		Views: engine,
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

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("healthy!")
	})

	api.LoadRoutes(app)

	app.Static("/", "./server/views/web/static", fiber.Static{
		Compress:      true,
		ByteRange:     true,
		Browse:        true,
		Index:         "home.html",
		CacheDuration: 1 * time.Second,
		MaxAge:        3600,
	})
	return app
}
