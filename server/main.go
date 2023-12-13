package server

import (
	"aoc/server/api"
	env "aoc/server/config"
	"aoc/server/db"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/template/html/v2"
	"html/template"
	"log"
	"os"
	"os/signal"
	"strings"
	"time"
)

func Start() {
	env.LoadConfig()
	db.PrepareDatabase()

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

	log.Fatal(app.Listen(env.Config.AppHost))
}

// fiberApp calls loadRoutes internally
func fiberApp() *fiber.App {
	engine := html.New("./server/views/web", ".html")
	if strings.HasPrefix(env.Config.AppHost, "localhost") {
		engine.Reload(true)
	}

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

	app.Use(func(c *fiber.Ctx) error {
		ip := c.Get("X-Real-IP")
		if ip == "" {
			ip = c.Get("X-Forwarded-For")
		}
		if ip == "" {
			ip = c.IP()
		}

		c.Locals("realIP", ip)
		return c.Next()
	})

	app.Use(logger.New(logger.Config{
		Format: "${time} [${locals:realIP}]:${port} ${pid} ${locals:requestid} ${status} - ${latency} ${method} ${path} ?${queryParams}\n\t\t\t\t\tUser-Agent: ${header:User-Agent}\n",
	}))

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("healthy!")
	})
	app.Use(recover.New())

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
