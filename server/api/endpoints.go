package api

import (
	"aoc/server/db"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"strings"
	"time"
)

type RequestData struct {
	Day      int    `json:"day"`
	Part     int    `json:"part"`
	Input    string `json:"input"`
	Solution int    `json:"solution"`
}

var Limiter = limiter.New(limiter.Config{
	Max:        6,
	Expiration: 1 * time.Minute,
	LimitReached: func(ctx *fiber.Ctx) error {
		return ctx.Status(429).SendString("Too Many Requests. Wait a minute!")
	},
})

func LoadRoutes(app *fiber.App) {
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Render("list-problems", fiber.Map{
			"Problems": ListProblems(),
		})
	})

	app.Get("/solve/day/:day/part/:part", func(ctx *fiber.Ctx) error {
		day, err := ctx.ParamsInt("day")
		if err != nil {
			return err
		}
		part, err := ctx.ParamsInt("part")
		if err != nil {
			return err
		}
		if day <= 0 {
			return fmt.Errorf("invalid day")
		}
		if part <= 0 {
			return fmt.Errorf("invalid part")
		}

		return ctx.Render("solve-problem", fiber.Map{
			"Day":  day,
			"Part": part,
		})
	})

	app.Post("/check-solution", Limiter, func(ctx *fiber.Ctx) error {
		realIp, ok := ctx.Locals("realIP").(string)
		fmt.Println("realIp", realIp)
		rl := &db.RequestLog{
			CreatedAt: time.Now(),
			IP:        ctx.IP(),
		}
		if ok && realIp != "" {
			rl.IP = realIp
		}
		_ = rl.Insert(ctx.Context())

		data := &RequestData{}
		err := ctx.BodyParser(data)
		if err != nil {
			fmt.Println(err)
			return err
		}
		rl.Day = data.Day
		rl.Part = data.Part
		_ = rl.Update(ctx.Context())

		if strings.TrimSpace(data.Input) == "" {
			return ctx.Status(400).SendString("Your input is empty!")
		}
		if data.Day <= 0 {
			return ctx.Status(400).SendString("Day can not be less than 0!")
		}
		if data.Part <= 0 {
			return ctx.Status(400).SendString("Part can not be less than 0!")
		}

		for _, problem := range ListProblems() {
			if problem.Day == data.Day && problem.Part == data.Part {
				rl.Valid = true

				solution, err := problem.Handler(strings.NewReader(data.Input))
				if err != nil {
					return err
				}

				fmt.Println(solution, " -> ", data.Solution)
				if solution == data.Solution {
					rl.CorrectAnswer = true
					_ = rl.Update(ctx.Context())
					return ctx.Status(200).SendString("Your answer is right!")
				}
				rl.CorrectAnswer = false
				_ = rl.Update(ctx.Context())
				return ctx.Status(400).SendString("Your answer is WRONG!")
			}
		}
		rl.Message = "Problem not found."
		rl.Update(ctx.Context())
		return ctx.Status(404).SendString("Problem not found.")
	})
}
