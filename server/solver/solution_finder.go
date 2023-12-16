package solver

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"strings"
	"time"
)

type SolutionFinder struct {
	ResultLine string
}

func (s *SolutionFinder) Close() error {
	fmt.Println("close", time.Now())
	return nil
}

func (s *SolutionFinder) Write(byteSlice []byte) (n int, err error) {
	scanner := bufio.NewScanner(bytes.NewReader(byteSlice))
	for scanner.Scan() {
		line := scanner.Text()
		log.Debug("child output: ", "[", line, "]")
		if strings.HasPrefix(line, "result: ") {
			s.ResultLine = line
		}
	}

	return len(byteSlice), scanner.Err()
}
