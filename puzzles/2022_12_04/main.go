package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	content, err := os.Open("2022_04/dataset.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(content)

	var areas []string
	for scanner.Scan() {
		areas = append(areas, scanner.Text())
	}

	overlapCountPart1 := 0
	overlapCountPart2 := 0
	for _, area := range areas {
		workers := strings.Split(area, ",")

		worker1Area := strings.Split(workers[0], "-")
		worker2Area := strings.Split(workers[1], "-")

		worker1AreaStart, _ := strconv.Atoi(worker1Area[0])
		worker1AreaEnd, _ := strconv.Atoi(worker1Area[1])
		worker2AreaStart, _ := strconv.Atoi(worker2Area[0])
		worker2AreaEnd, _ := strconv.Atoi(worker2Area[1])

		if worker1AreaStart >= worker2AreaStart && worker1AreaEnd <= worker2AreaEnd || worker2AreaStart >= worker1AreaStart && worker2AreaEnd <= worker1AreaEnd {
			overlapCountPart1++
		}

		if worker1AreaStart <= worker2AreaStart && worker1AreaEnd >= worker2AreaStart || worker2AreaStart <= worker1AreaStart && worker2AreaEnd >= worker1AreaStart {
			overlapCountPart2++
		}

	}
	println("Result Part 1:", overlapCountPart1)
	println("Result Part 2:", overlapCountPart2)

}
