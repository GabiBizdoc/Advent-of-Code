package solution

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseLine(line string) ([]int, error) {
	numsLiteral := strings.Split(line, " ")
	nums := make([]int, 0, len(numsLiteral))
	for _, n := range numsLiteral {
		n = strings.TrimSpace(n)
		if n != "" {
			if x, err := strconv.Atoi(n); err != nil {
				return nil, err
			} else {
				nums = append(nums, x)
			}
		}
	}
	return nums, nil
}

func computeNextLine(nums []int) []int {
	next := make([]int, 0, len(nums)-1)
	for i := 1; i < len(nums); i++ {
		next = append(next, nums[i]-nums[i-1])
	}
	return next
}

func isZero(nums []int) bool {
	for _, num := range nums {
		if num != 0 {
			return false
		}
	}
	return true
}

func findLast(nums []int) int {
	lastNums := make([]int, 0)
	for isZero(nums) == false {
		lastNums = append(lastNums, nums[0])
		nums = computeNextLine(nums)
	}
	t1 := 0
	for i := len(lastNums) - 1; i >= 0; i-- {
		t2 := lastNums[i]
		t1 = t2 - t1
	}
	return t1
}

func solveChallenge(inputFilePath string) (int, error) {
	fmt.Println(inputFilePath)
	solution := 0

	file, err := os.Open(inputFilePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		nums, err := parseLine(line)
		if err != nil {
			return 0, err
		}

		sol := findLast(nums)
		solution += sol
	}

	if err = scanner.Err(); err != nil {
		return 0, err
	}

	fmt.Println("solution: ", solution)
	return solution, nil
}

func SolveChallenge(inputFilePath string) (int, error) {
	return solveChallenge(inputFilePath)
}
