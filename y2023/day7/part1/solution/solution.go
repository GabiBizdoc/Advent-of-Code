package solution

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Solve(file io.Reader) (int, error) {
	var solution int

	hands, err := readData(file)
	if err != nil {
		return 0, err
	}

	slices.SortFunc(hands, func(a, b *Hand) int {
		return CmpHands(a, b)
	})

	slices.Reverse(hands)

	for i, hand := range hands {
		solution += hand.Bid * (i + 1)
	}

	fmt.Println(solution)
	return solution, nil
}

func solveChallenge(inputFilePath string) (int, error) {
	file, err := os.Open(inputFilePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	return Solve(file)
}

func readData(file io.Reader) ([]*Hand, error) {
	data := make([]*Hand, 0, 50)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		parts := strings.Split(line, " ")
		bid, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, err
		}
		data = append(data, NewHand(parts[0], bid))
	}

	return data, scanner.Err()
}

func SolveChallenge(inputFilePath string) (int, error) {
	return solveChallenge(inputFilePath)
}
