package solution

import (
	"fmt"
	"sync"
)

func SolveChallenge(inputFilePath string) (int, error) {
	return solveChallengeParallel(inputFilePath)
}

// seed -> soil -> fertilizer -> .. -> location
var almanacPath = []string{"seed", "soil", "fertilizer", "water", "light", "temperature", "humidity", "location"}

func solveChallengeParallel(inputFilePath string) (int, error) {
	var solution = NewSolution()
	almanac, err := readAlmanac(inputFilePath)
	if err != nil {
		return 0, err
	}

	var wg sync.WaitGroup
	for i := 0; i < len(almanac.Seeds); i += 2 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			var sol = NewSolution()
			start, seedRage := almanac.Seeds[i], almanac.Seeds[i+1]
			for j := 0; j < seedRage; j++ {
				seed := start + j
				next, err := almanac.GetPath(seed, almanacPath...)
				if err != nil {
					panic(err)
					//return 0, err
				}
				sol.UnsafeUpdate(next)
			}
			solution.SafeUpdate(sol.Value())
		}(i)
	}
	wg.Wait()

	fmt.Println(solution)
	return solution.Value(), nil
}

func solveChallengeSimple(inputFilePath string) (int, error) {
	var solution = NewSolution()
	almanac, err := readAlmanac(inputFilePath)
	if err != nil {
		return 0, err
	}

	for i := 0; i < len(almanac.Seeds); i += 2 {
		start, seedRage := almanac.Seeds[i], almanac.Seeds[i+1]

		for i := 0; i < seedRage; i++ {
			seed := start + i
			next, err := almanac.GetPath(seed, almanacPath...)
			if err != nil {
				panic(err)
				//return 0, err
			}
			solution.UnsafeUpdate(next)
		}
	}

	fmt.Println("solution:", solution)
	return solution.Value(), nil
}

func solveChallenge(inputFilePath string) (int, error) {
	var solution = NewSolution()
	almanac, err := readAlmanac(inputFilePath)
	if err != nil {
		return 0, err
	}

	for i := 0; i < len(almanac.Seeds); i += 2 {
		start, seedRage := almanac.Seeds[i], almanac.Seeds[i+1]

		for i := 0; i < seedRage; i++ {
			seed := start + i
			next, err := almanac.GetPath(seed, almanacPath...)
			if err != nil {
				panic(err)
				//return 0, err
			}
			solution.UnsafeUpdate(next)
		}
	}

	fmt.Println("solution:", solution)
	return solution.Value(), nil
}
