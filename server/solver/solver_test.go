package solver_test

import (
	env "aoc/server/config"
	"aoc/server/solver"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	fmt.Println(filename, dir)
	parentDir := filepath.Join(dir, "../../")
	err := os.Chdir(parentDir)
	if err != nil {
		panic(err)
	}

	env.Config.IsDev = true

	m.Run()
}

type ProblemData struct {
	Day, Part int
	Expected  int
}

func (p ProblemData) Input() string {
	data, err := os.ReadFile(fmt.Sprintf("./y2023/day_%02d/input-long.txt", p.Day))
	if err != nil {
		panic(err)
	}
	return string(data)
}

func TestSolveProblem(t *testing.T) {
	problems := []ProblemData{
		{1, 1, 54601},
		{1, 2, 54078},
		{2, 1, 2348},
		{2, 2, 76008},
		{3, 1, 543867},
		{3, 2, 79613331},
		{4, 1, 28750},
		{4, 2, 10212704},
		{5, 1, 261668924},
		{5, 2, 24261545},
		{6, 1, 1624896},
		{6, 2, 32583852},
		{7, 1, 249204891},
		{7, 2, 249666369},
		{8, 1, 15517},
		{8, 2, 14935034899483},
		{9, 1, 1798691765},
		{9, 2, 1104},
		//{10, 1, 7066},
		//{10, 2, 401},
		{11, 1, 10289334},
		{11, 2, 649862989626},
		{12, 1, 7260},
		{12, 2, 1909291258644},
		{13, 1, 29130},
		{13, 2, 33438},
		{14, 1, 106378},
		{14, 2, 90795},
		{15, 1, 517551},
		{15, 2, 286097},
		{16, 1, 7939},
		{16, 2, 8318},
		//{17, 1, 000},
		//{17, 2, 000},
		//{18, 1, 000},
		//{18, 2, 000},
		//{19, 1, 000},
		//{19, 2, 000},
		//{20, 1, 000},
		//{20, 2, 000},
		//{21, 1, 000},
		//{21, 2, 000},
		//{22, 1, 000},
		//{22, 2, 000},
		//{23, 1, 000},
		//{23, 2, 000},
		//{24, 1, 000},
		//{24, 2, 000},
		//{25, 1, 000},
		//{25, 2, 10212704},
	}

	var wg sync.WaitGroup
	wg.Add(len(problems))

	for i, data := range problems {
		go func() {
			defer wg.Done()
			result := solver.SolveProblem(30*time.Second, data.Day, data.Part, data.Input())
			if result.Err != nil {
				panic(result.Err)
			}
			if result.Solution != data.Expected {
				t.Errorf("test %d day_%d_%d failed. expected: %d got %d\n", i, data.Day, data.Part, data.Expected, result.Solution)
			}
		}()
	}
	wg.Wait()
}
