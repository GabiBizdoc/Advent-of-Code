package solution

import (
	"bufio"
	"fmt"
	"io"
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
	s := 0
	for isZero(nums) == false {
		s += nums[len(nums)-1]
		nums = computeNextLine(nums)
	}
	return s
}
func solveChallenge(inputFilePath string) (int, error) {
	fmt.Println(inputFilePath)
	file, err := os.Open(inputFilePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	return Solve(file)
}

func Solve(file io.Reader) (int, error) {
	solution := 0
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

	return solution, scanner.Err()
}

func SolveChallenge(inputFilePath string) (int, error) {
	return solveChallenge(inputFilePath)
}
