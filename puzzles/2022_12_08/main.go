package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func main() {
	defer timeTrack(time.Now(), "Calculation")
	content, err := os.Open("puzzles/2022_12_08/dataset.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(content)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var mapOfTree = make(map[int]map[int]int)
	for y, line := range lines {
		if line == "" {
			continue
		}
		for x, character := range line {
			value, err := strconv.Atoi(string(character))
			if err != nil {
				log.Fatal(err)
			}
			if _, ok := mapOfTree[x]; !ok {
				mapOfTree[x] = make(map[int]int)
			}
			mapOfTree[x][y] = value
		}
	}
	treeCount := 0
	scoreView := 0
	for x, line := range mapOfTree {
		for y, height := range line {

			var isNextTreeRightSmaller, isNextTreeLeftSmaller, isNextTreeUpSmaller, isNextTreeDownSmaller bool
			var treeScore, countRight, countLeft, countUp, countDown int

			isNextTreeRightSmaller, countRight = isNextTreeTaller(x, y, "right", x, y, mapOfTree, height, 0)
			isNextTreeLeftSmaller, countLeft = isNextTreeTaller(x, y, "left", x, y, mapOfTree, height, 0)
			isNextTreeUpSmaller, countUp = isNextTreeTaller(x, y, "up", x, y, mapOfTree, height, 0)
			isNextTreeDownSmaller, countDown = isNextTreeTaller(x, y, "down", x, y, mapOfTree, height, 0)

			treeScore = countRight * countLeft * countUp * countDown
			if treeScore > scoreView {
				scoreView = treeScore
			}

			if !isNextTreeRightSmaller || !isNextTreeLeftSmaller || !isNextTreeUpSmaller || !isNextTreeDownSmaller {
				treeCount++
				continue
			}

		}
	}

	log.Println("Visible trees", treeCount)
	log.Println("Score view", scoreView)

}

func isNextTreeTaller(x int, y int, direction string, nextX int, nextY int, mapOfTree map[int]map[int]int, highestPoint int, viewTreeCount int) (bool, int) {

	currentTreeHeight := mapOfTree[x][y]

	switch direction {
	case "right":
		nextX++
	case "left":
		nextX--
	case "up":
		nextY--
	case "down":
		nextY++
	}

	if nextX < 0 || nextY < 0 {
		return false, viewTreeCount
	}

	if value, ok := mapOfTree[nextX]; !ok {
		return false, viewTreeCount
	} else if _, ok := value[nextY]; !ok {
		return false, viewTreeCount
	}

	viewTreeCount++

	nextHeight := mapOfTree[nextX][nextY]

	if currentTreeHeight > highestPoint {
		highestPoint = currentTreeHeight
	}

	if highestPoint <= nextHeight {
		return true, viewTreeCount
	} else {
		return isNextTreeTaller(x, y, direction, nextX, nextY, mapOfTree, highestPoint, viewTreeCount)
	}

}
