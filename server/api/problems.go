package api

import (
	"io"

	day1part1 "aoc/y2023/day_01/part1/solution"
	day1part2 "aoc/y2023/day_01/part2/solution"

	day2part1 "aoc/y2023/day_02/part1/solution"
	day2part2 "aoc/y2023/day_02/part2/solution"

	day3part1 "aoc/y2023/day_03/part1/solution"
	day3part2 "aoc/y2023/day_03/part2/solution"

	day4part1 "aoc/y2023/day_04/part1/solution"
	day4part2 "aoc/y2023/day_04/part2/solution"

	day5part1 "aoc/y2023/day_05/part1_optimal/solution"
	day5part2 "aoc/y2023/day_05/part2_optimal/solution"

	day6part1 "aoc/y2023/day_06/part1/solution"
	day6part2 "aoc/y2023/day_06/part2/solution"

	day7part1 "aoc/y2023/day_07/part1/solution"
	day7part2 "aoc/y2023/day_07/part2/solution"

	day8part1 "aoc/y2023/day_08/part1/solution"
	day8part2 "aoc/y2023/day_08/part2/solution"

	day9part1 "aoc/y2023/day_09/part1/solution"
	day9part2 "aoc/y2023/day_09/part2/solution"

	//day10part1 "aoc/y2023/day_10/part1/solution"
	//day10part2 "aoc/y2023/day_10/part2/solution"

	day11part1 "aoc/y2023/day_11/part1/solution"
	day11part2 "aoc/y2023/day_11/part2/solution"

	day12part1 "aoc/y2023/day_12/part1/solution"
	day12part2 "aoc/y2023/day_12/part2/solution"

	day13part1 "aoc/y2023/day_13/part1/solution"
	day13part2 "aoc/y2023/day_13/part2/solution"

	day14part1 "aoc/y2023/day_14/part1/solution"
	day14part2 "aoc/y2023/day_14/part2_brute_force/solution"

	day15part1 "aoc/y2023/day_15/part1/solution"
	day15part2 "aoc/y2023/day_15/part2/solution"

	day16part1 "aoc/y2023/day_16/part1/solution"
	day16part2 "aoc/y2023/day_16/part2/solution"
	//day17part1 "aoc/y2023/day_17/part1/solution"
	//day17part2 "aoc/y2023/day_17/part2/solution"
	//day18part1 "aoc/y2023/day_18/part1/solution"
	//day18part2 "aoc/y2023/day_18/part2/solution"
	//day19part1 "aoc/y2023/day_19/part1/solution"
	//day19part2 "aoc/y2023/day_19/part2/solution"
	//day20part1 "aoc/y2023/day_20/part1/solution"
	//day20part2 "aoc/y2023/day_20/part2/solution"
	//day21part1 "aoc/y2023/day_21/part1/solution"
	//day21part2 "aoc/y2023/day_21/part2/solution"
	//day22part1 "aoc/y2023/day_22/part1/solution"
	//day22part2 "aoc/y2023/day_22/part2/solution"
	//day23part1 "aoc/y2023/day_23/part1/solution"
	//day23part2 "aoc/y2023/day_23/part2/solution"
	//day24part1 "aoc/y2023/day_24/part1/solution"
	//day24part2 "aoc/y2023/day_24/part2/solution"
	//day25part1 "aoc/y2023/day_25/part1/solution"
	//day25part2 "aoc/y2023/day_25/part2/solution"
)

type ProblemHandler func(i io.Reader) (int, error)

type Problem struct {
	Id      int
	Name    string
	Day     int
	Part    int
	Handler ProblemHandler
}

func NewProblem(name string, day int, part int, handler ProblemHandler) *Problem {
	return &Problem{Name: name, Day: day, Part: part, Handler: handler}
}

func ListProblems() []*Problem {
	problems := make([]*Problem, 0)
	problems = append(problems, NewProblem("Trebuchet?!", 1, 1, day1part1.Solve))
	problems = append(problems, NewProblem("Trebuchet?!", 1, 2, day1part2.Solve))

	problems = append(problems, NewProblem("Cube Conundrum", 2, 1, day2part1.Solve))
	problems = append(problems, NewProblem("Cube Conundrum", 2, 2, day2part2.Solve))

	problems = append(problems, NewProblem("Gear Ratios", 3, 1, day3part1.Solve))
	problems = append(problems, NewProblem("Gear Ratios", 3, 2, day3part2.Solve))

	problems = append(problems, NewProblem("Scratchcards", 4, 1, day4part1.Solve))
	problems = append(problems, NewProblem("Scratchcards", 4, 2, day4part2.Solve))

	problems = append(problems, NewProblem("If You Give A Seed A Fertilizer", 5, 1, day5part1.Solve))
	problems = append(problems, NewProblem("If You Give A Seed A Fertilizer", 5, 2, day5part2.Solve))

	problems = append(problems, NewProblem("Wait For It", 6, 1, day6part1.Solve))
	problems = append(problems, NewProblem("Wait For It", 6, 2, day6part2.Solve))

	problems = append(problems, NewProblem("Camel Cards", 7, 1, day7part1.Solve))
	problems = append(problems, NewProblem("Camel Cards", 7, 2, day7part2.Solve))

	problems = append(problems, NewProblem("Haunted Wasteland", 8, 1, day8part1.Solve))
	problems = append(problems, NewProblem("Haunted Wasteland", 8, 2, day8part2.Solve))

	problems = append(problems, NewProblem("Mirage Maintenance", 9, 1, day9part1.Solve))
	problems = append(problems, NewProblem("Mirage Maintenance", 9, 2, day9part2.Solve))

	//problems = append(problems, NewProblem("Pipe Maze", 10, 1, day10part1.Solve))
	//problems = append(problems, NewProblem("Pipe Maze", 10, 2, day10part2.Solve))

	problems = append(problems, NewProblem("Cosmic Expansion", 11, 1, day11part1.Solve))
	problems = append(problems, NewProblem("Cosmic Expansion", 11, 2, day11part2.Solve))

	problems = append(problems, NewProblem("Hot Springs", 12, 1, day12part1.Solve2))
	problems = append(problems, NewProblem("Hot Springs", 12, 2, day12part2.Solve2))

	problems = append(problems, NewProblem("Point of Incidence", 13, 1, day13part1.Solve))
	problems = append(problems, NewProblem("Point of Incidence", 13, 2, day13part2.Solve))

	problems = append(problems, NewProblem("Parabolic Reflector Dish", 14, 1, day14part1.Solve))
	problems = append(problems, NewProblem("Parabolic Reflector Dish", 14, 2, day14part2.Solve))

	problems = append(problems, NewProblem("Lens Library", 15, 1, day15part1.Solve))
	problems = append(problems, NewProblem("Lens Library", 15, 2, day15part2.Solve))

	problems = append(problems, NewProblem("The Floor Will Be Lava", 16, 1, day16part1.Solve))
	problems = append(problems, NewProblem("The Floor Will Be Lava", 16, 2, day16part2.Solve))

	//problems = append(problems, NewProblem("Clumsy Crucible", 17, 1, day17part1.Solve2))
	//problems = append(problems, NewProblem("Clumsy Crucible", 17, 2, day17part2.Solve2))
	for i, problem := range problems {
		problem.Id = i + 1
	}
	return problems
}
