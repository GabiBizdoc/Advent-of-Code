package main

import (
	"aoc/server"
	"aoc/server/api"
	"errors"
	"flag"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
	"os"
	"runtime/debug"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	shouldStartServer := flag.Bool("server", false, "server")
	dayStr := flag.String("day", "1", "target day")
	partStr := flag.String("part", "2", "target part")
	flag.Parse()

	if *shouldStartServer {
		startWebserver()
	} else {
		result, err := solveProblem(dayStr, partStr)
		if err != nil {
			log.Error(err)
			panic(err)
		}
		fmt.Println("result:", result, "took:", time.Since(start))
	}
}

type Solution struct {
	Solution int
	Err      error
}

func solveProblem(dayStr, partStr *string) (int, error) {
	day, err := strconv.Atoi(*dayStr)
	if err != nil {
		panic(err)
	}

	part, err := strconv.Atoi(*partStr)
	if err != nil {
		panic(err)
	}
	if day <= 0 {
		panic("invalid day")
	}
	if part <= 0 {
		panic("invalid part")
	}

	log.Info("day:", day, "part:", part)
	for _, problem := range api.ListProblems() {
		if problem.Day == day && problem.Part == part {
			log.Info("problem found for day:", day, " part:", part)

			solution := func(handler api.ProblemHandler) *Solution {
				solution := &Solution{}
				defer func() {
					if r := recover(); r != nil {
						debug.PrintStack()
						fmt.Println("Recovered in f", r)
						solution.Err = fmt.Errorf("failed to run handler")
					}
				}()
				solution.Solution, solution.Err = problem.Handler(os.Stdin)
				return solution
			}(problem.Handler)

			return solution.Solution, solution.Err
		}
	}
	return 0, errors.New("challenge not found")
}

func startWebserver() {
	godotenv.Load()
	server.Start()
}
