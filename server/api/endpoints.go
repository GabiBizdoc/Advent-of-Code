package api

import (
	"aoc/server/db"
	"aoc/server/solver"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"strings"
	"time"
)

type RequestData struct {
	Day      int         `json:"day" validate:"required"`
	Part     int         `json:"part" validate:"required"`
	Input    string      `json:"input" validate:"required"`
	Solution json.Number `json:"solution" validate:"required"`
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

	app.Post("/check-solution", Limiter, checkSolutionHandler)
}

func checkSolutionHandler(ctx *fiber.Ctx) error {
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
	defer func(rl *db.RequestLog, ctx context.Context) {
		err := rl.Update(ctx)
		if err != nil {
			fmt.Println(err)
		}
	}(rl, ctx.Context())

	data := &RequestData{}
	err := ctx.BodyParser(data)
	//fmt.Println(data, ctx.FormValue("solution"))
	if err != nil {
		fmt.Println(err)
		return err
	}
	rl.Day = data.Day
	rl.Part = data.Part

	if strings.TrimSpace(data.Input) == "" {
		return ctx.Status(400).SendString("Your input is empty!")
	}
	if data.Day <= 0 {
		return ctx.Status(400).SendString("Day can not be less than 0!")
	}
	if data.Part <= 0 {
		return ctx.Status(400).SendString("Part can not be less than 0!")
	}
	userSolution, err := data.Solution.Int64()
	if err != nil {
		return err
	}

	for _, problem := range ListProblems() {
		if problem.Day == data.Day && problem.Part == data.Part {
			rl.Valid = true
			result := solver.SolveProblem(30*time.Second, data.Day, data.Part, data.Input)
			if result.Err != nil {
				return result.Err
			}
			fmt.Println(
				"day", problem.Day,
				"part", problem.Part,
				"solution: ", result.Solution,
				"execution_time: ", result.ExecutionTime,
				"total_time: ", result.RealTime,
				"user's solution: ", data.Solution)
			if int64(result.Solution) == userSolution {
				rl.CorrectAnswer = true
				return ctx.Status(200).SendString("Your answer is right!")
			}
			rl.CorrectAnswer = false
			return ctx.Status(400).SendString("Your answer is WRONG!")
		}
	}
	rl.Message = "Problem not found."
	return ctx.Status(404).SendString("Problem not found.")
}
