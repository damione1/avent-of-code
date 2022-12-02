package main

import (
	"bufio"
	"fmt"
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

	/* dataset example
	   2-4,6-8
	   2-3,4-5
	   5-7,7-9
	   2-8,3-7
	   6-6,4-6
	   2-6,4-8
	*/
	overlapCount := 0
	for _, area := range areas {
		//split line at the comma
		workers := strings.Split(area, ",")

		worker1Area := strings.Split(workers[0], "-")
		worker2Area := strings.Split(workers[1], "-")

		//convert strings to ints
		worker1AreaStart, _ := strconv.Atoi(worker1Area[0])
		worker1AreaEnd, _ := strconv.Atoi(worker1Area[1])
		worker2AreaStart, _ := strconv.Atoi(worker2Area[0])
		worker2AreaEnd, _ := strconv.Atoi(worker2Area[1])

		//find if one worker area is completly in the other worker's area
		if worker1AreaStart >= worker2AreaStart && worker1AreaEnd <= worker2AreaEnd {
			overlapCount++
			//print the area that is completly in the other worker's area
			fmt.Println("Worker 1 area is completly in worker 2 area:", area)
		} else if worker2AreaStart >= worker1AreaStart && worker2AreaEnd <= worker1AreaEnd {
			overlapCount++
			//print the area that is completly in the other worker's area
			fmt.Println("Worker 2 area is completly in worker 1 area:", area)
		} else {
			fmt.Println("No overlap:", area)
		}

	}
	println(overlapCount)

}
