package solution

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func solveChallenge(inputFilePath string) (int, error) {
	var solution int

	hands, err := readData(inputFilePath)
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

func readData(inputFilePath string) ([]*Hand, error) {
	data := make([]*Hand, 0, 50)

	file, err := os.Open(inputFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
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

	if err = scanner.Err(); err != nil {
		return nil, err
	}

	return data, nil
}

func SolveChallenge(inputFilePath string) (int, error) {
	return solveChallenge(inputFilePath)
}
