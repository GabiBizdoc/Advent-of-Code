package solution

import (
	testcom "aoc/com/test_com"
	"fmt"
	"os"
	"reflect"
	"testing"
)

// IMPORTANT NOTE:
// These tests were written with the assumption that the interval notation used is [start, end].
// However, it has been discovered that the actual interval notation used is [start, end).
//
// As a result, all tests may produce incorrect results and should not be relied upon.

func TestMain(m *testing.M) {
	//DebugEnabled = true
	m.Run()
}

func testInterval(t *testing.T, a *Map, iStart, iEnd int, expected [][2]int) {
	response := a.nextInterval(iStart, iEnd)
	if !reflect.DeepEqual(response, expected) {
		t.Errorf("FAIL: for %v expected %v got %v", [...]int{iStart, iEnd}, expected, response)
	}
	total := 0
	for _, i := range response {
		intervalRange := i[1] - i[0]
		t.Log("intervalRange", i, intervalRange)

		if intervalRange < 0 {
			t.Errorf("FAIL: negative range for %v, => %d", [...]int{iStart, iEnd}, intervalRange)
		}
		total += intervalRange
	}
	expectedSum := iEnd - iStart - len(response) + 1
	if total != expectedSum {
		t.Errorf("FAIL: wrong sum for interval for %v, expected %d got %d ", [...]int{iStart, iEnd}, expectedSum, total)
	}
}
func TestMap_NextInterval222(t *testing.T) {
	a := NewMap("", "")
	a.AddRow(1000, 0, 500)
	a.AddRow(0, 1000, 500)

	testInterval(t, a, 10, 2000, [][2]int{{1010, 1500}, {501, 999}, {0, 500}, {1501, 2000}})
}

func TestMap_NextInterval(t *testing.T) {
	a := NewMap("", "")
	a.AddRow(100, 200, 100)
	a.AddRow(1000, 2000, 1000)

	testInterval(t, a, 190, 500, [][2]int{{290, 300}, {201, 500}})
	testInterval(t, a, 0, 5000, [][2]int{ // pass end
		{0, 99}, {200, 300}, {201, 999}, {2000, 3000}, {2001, 5000},
	})
	testInterval(t, a, 10, 20, [][2]int{{10, 20}})              // before intervals
	testInterval(t, a, 3000, 4000, [][2]int{{3000, 4000}})      // after intervals
	testInterval(t, a, 111, 122, [][2]int{{211, 222}})          // in the middle
	testInterval(t, a, 90, 110, [][2]int{{90, 99}, {200, 210}}) // pass start
}

func TestMap_NextInterval22(t *testing.T) {
	t.Skip("this test is only for visualization purposes")

	f, err := os.Open("../" + testcom.Part2ShortFilepath)
	if err != nil {
		panic(err)
	}

	_, a, err := readAlmanac(f)
	if err != nil {
		panic(err)
	}

	almanacPath := []string{"seed", "soil", "fertilizer", "water", "light", "temperature", "humidity", "location"}
	ap, err := a.GetPath(almanacPath...)
	if err != nil {
		panic(err)
	}

	for i := 79; i < 93; i++ {
		fmt.Println("i = ", i)
		fmt.Println(ap.Traverse(i))
	}
	fmt.Println(ap.TraverseInterval(79, 84))
}

func TestMap_NextInterval2(t *testing.T) {
	t.Skip("this test is only for visualization purposes")

	a := NewMap("", "")

	a.AddRow(53, 49, 8)
	a.AddRow(11, 0, 42)
	a.AddRow(0, 42, 7)
	a.AddRow(7, 57, 4)

	for _, row := range a.Rows {
		line := make([]rune, 100)
		for i := 0; i < len(line); i++ {
			if i >= row.SourceStart() && i <= row.SourceEnd() {
				line[i] = 'x'
			}
		}

		line2 := make([]rune, len(line))
		for i := 0; i < len(line2); i++ {
			if i >= row.DestinationStart() && i <= row.DestinationEnd() {
				line2[i] = 'y'
			}
		}
		fmt.Println("source", row.SourceStart(), row.SourceEnd())
		fmt.Println(string(line))
		fmt.Println("destin", row.DestinationStart(), row.DestinationEnd())
		fmt.Println(string(line2))
	}

	testInterval(t, a, 58, 69, [][2]int{{0, 0}})
}
