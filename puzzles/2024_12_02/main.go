package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// iota direction
const (
	Up = iota
	Down
)

func main() {
	content, err := os.Open("puzzles/2024_12_02/dataset.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(content)

	var reports [][]int
	var solutionOne int
	var solutionTwo int
	for scanner.Scan() {
		//lines = append(lines, scanner.Text())
		var report []int
		for _, v := range strings.Split(scanner.Text(), " ") {
			i, _ := strconv.Atoi(v)
			report = append(report, i)
		}
		reports = append(reports, report)
		test1, test2 := isReportSafe(report)

		if test1 {
			solutionOne++
		}

		if test2 {
			solutionTwo++
		}

	}

	fmt.Println(fmt.Sprintf("Solution One: %v", solutionOne))
	fmt.Println(fmt.Sprintf("Solution Two: %v", solutionTwo))

}

func isReportSafe(reports []int) (bool, bool) {
	if len(reports) < 2 {
		return true, true
	}

	if checkSafety(reports) {
		return true, true
	}

	for i := range reports {
		withoutCurrent := make([]int, 0, len(reports)-1)
		withoutCurrent = append(withoutCurrent, reports[:i]...)
		withoutCurrent = append(withoutCurrent, reports[i+1:]...)

		if checkSafety(withoutCurrent) {
			return false, true
		}
	}

	return false, false
}

func checkSafety(reports []int) bool {
	if len(reports) < 2 {
		return true
	}

	direction := reports[1] - reports[0]
	if direction == 0 || abs(direction) > 3 {
		return false
	}

	for i := 2; i < len(reports); i++ {
		diff := reports[i] - reports[i-1]
		if diff == 0 || abs(diff) > 3 || (direction > 0 && diff < 0) || (direction < 0 && diff > 0) {
			return false
		}
	}

	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
