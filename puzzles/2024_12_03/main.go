package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	content, err := os.Open("puzzles/2024_12_03/dataset.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(content)

	var instructions string

	for scanner.Scan() {
		instructions = instructions + scanner.Text()
	}

	re1 := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches1 := re1.FindAllStringSubmatch(instructions, -1)

	partOneSum := 0
	for _, match := range matches1 {
		num1 := match[1]
		num2 := match[2]

		num1Int := 0
		num2Int := 0

		fmt.Sscanf(num1, "%d", &num1Int)
		fmt.Sscanf(num2, "%d", &num2Int)

		partOneSum += num1Int * num2Int
	}

	fmt.Println("Part One Sum: ", partOneSum)

}
