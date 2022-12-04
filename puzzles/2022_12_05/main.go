package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	content, err := os.Open("puzzles/2022_12_05/dataset.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(content)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	//	stacksSchema := `
	//[M]                     [N] [Z]
	//[F]             [R] [Z] [C] [C]
	//[C]     [V]     [L] [N] [G] [V]
	//[W]     [L]     [T] [H] [V] [F] [H]
	//[T]     [T] [W] [F] [B] [P] [J] [L]
	//[D] [L] [H] [J] [C] [G] [S] [R] [M]
	//[L] [B] [C] [P] [S] [D] [M] [Q] [P]
	//[B] [N] [J] [S] [Z] [W] [F] [W] [R]
	// 1   2   3   4   5   6   7   8   9
	// `
	stacks := make([][]string, 10)
	stacks[0] = []string{}
	stacks[1] = []string{"B", "L", "D", "T", "W", "C", "F", "M"}
	stacks[2] = []string{"N", "B", "L"}
	stacks[3] = []string{"J", "C", "H", "T", "L", "V"}
	stacks[4] = []string{"S", "P", "J", "W"}
	stacks[5] = []string{"Z", "S", "C", "F", "T", "L", "R"}
	stacks[6] = []string{"W", "D", "G", "B", "H", "N", "Z"}
	stacks[7] = []string{"F", "M", "S", "P", "V", "G", "C", "N"}
	stacks[8] = []string{"W", "Q", "R", "J", "F", "V", "C", "Z"}
	stacks[9] = []string{"R", "P", "M", "L", "H"}

	for stackNumber, stack := range stacks {
		fmt.Println(stackNumber, " ", stack)
	}

	for _, line := range lines {
		regex := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
		matches := regex.FindStringSubmatch(line)
		if len(matches) != 4 {
			log.Fatal("Wrong line format: ", line)
		}
		quantityToMove, _ := strconv.Atoi(matches[1])
		stackFrom, _ := strconv.Atoi(matches[2])
		stackTo, _ := strconv.Atoi(matches[3])

		/* Part 1
		for i := 0; i < quantityToMove; i++ {
			stacks[stackTo] = append(stacks[stackTo], stacks[stackFrom][len(stacks[stackFrom])-1])
			stacks[stackFrom] = stacks[stackFrom][:len(stacks[stackFrom])-1]
		}
		*/

		// Part 2
		stacks[stackTo] = append(stacks[stackTo], stacks[stackFrom][len(stacks[stackFrom])-quantityToMove:]...)
		stacks[stackFrom] = stacks[stackFrom][:len(stacks[stackFrom])-quantityToMove]

	}

	for stackNumber, stack := range stacks {
		fmt.Println(stackNumber, " ", stack)
	}

	for _, stack := range stacks {
		if len(stack) == 0 {
			continue
		}
		fmt.Print(stack[len(stack)-1], "")
	}

}
