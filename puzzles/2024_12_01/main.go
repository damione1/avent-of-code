package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	content, err := os.Open("puzzles/2024_12_01/dataset.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(content)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var list1 = make([]int, len(lines))
	var list2 = make([]int, len(lines))

	similarValuesMap := make(map[int]int)

	for i, line := range lines {
		parts := strings.Split(line, "   ")

		val1, _ := strconv.Atoi(parts[0])
		list1[i] = val1

		val2, _ := strconv.Atoi(parts[1])
		list2[i] = val2

		similarValuesMap[val1] = 0
	}

	sort.Ints(list1)
	sort.Ints(list2)

	var solutionOne int
	var solutionTwo int
	for i := 0; i < len(list1); i++ {
		if list1[i] > list2[i] {
			solutionOne += (list1[i] - list2[i])
		} else {
			solutionOne += (list2[i] - list1[i])
		}
		if _, ok := similarValuesMap[list2[i]]; ok {
			similarValuesMap[list2[i]]++
		}
	}

	for key, value := range similarValuesMap {
		solutionTwo += (key * value)
	}

	fmt.Println(fmt.Sprintf("Solution 1: %d", solutionOne))
	fmt.Println(fmt.Sprintf("Solution 2: %d", solutionTwo))
}
