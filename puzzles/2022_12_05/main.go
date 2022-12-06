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
	//split instructions and stack schema
	var instructions []string
	var stackSchema []string
	isInstructions := false
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			isInstructions = true
			continue
		}
		if isInstructions {
			instructions = append(instructions, line)
		} else {
			stackSchema = append([]string{line}, stackSchema...)
		}
	}

	//create stacks indexes
	index := make(map[int]int)
	for charIndex, character := range stackSchema[0] {
		if character != ' ' {
			stackNumber, _ := strconv.Atoi(string(character))
			index[stackNumber] = charIndex
		}
	}

	//create stacks
	stacksPart1 := make([][]string, len(index))
	stacksPart2 := make([][]string, len(index))
	for lineNumber, stackLine := range stackSchema {
		if lineNumber == 0 {
			continue
		}
		for stackNumber, charIndex := range index {
			if charIndex < len(stackLine) && stackLine[charIndex] != ' ' {
				stacksPart1[stackNumber-1] = append(stacksPart1[stackNumber-1], string(stackLine[charIndex]))
				stacksPart2[stackNumber-1] = append(stacksPart2[stackNumber-1], string(stackLine[charIndex]))
			}
		}
	}

	fmt.Println("Original stacks:")
	for stackNumber, stack := range stacksPart1 {
		fmt.Println(stackNumber, stack)
	}

	//process instructions
	for _, instruction := range instructions {
		regex := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
		matches := regex.FindStringSubmatch(instruction)
		if len(matches) != 4 {
			log.Fatal("Wrong line format: ", instruction)
		}
		quantityToMove, _ := strconv.Atoi(matches[1])
		stackFrom, _ := strconv.Atoi(matches[2])
		stackTo, _ := strconv.Atoi(matches[3])

		stackFrom--
		stackTo--

		//Part 1
		for i := 0; i < quantityToMove; i++ {
			stacksPart1[stackTo] = append(stacksPart1[stackTo], stacksPart1[stackFrom][len(stacksPart1[stackFrom])-1])
			stacksPart1[stackFrom] = stacksPart1[stackFrom][:len(stacksPart1[stackFrom])-1]
		}

		// Part 2
		stacksPart2[stackTo] = append(stacksPart2[stackTo], stacksPart2[stackFrom][len(stacksPart2[stackFrom])-quantityToMove:]...)
		stacksPart2[stackFrom] = stacksPart2[stackFrom][:len(stacksPart2[stackFrom])-quantityToMove]

	}

	fmt.Println("\n\n-------\n\nSchema part 1: ")
	responsePart1 := ""
	for stackNumber, stack := range stacksPart1 {
		fmt.Println(stackNumber, " ", stack)
		responsePart1 += stack[len(stack)-1]
	}
	fmt.Println("\nResponse: " + responsePart1)

	fmt.Println("\nSchema part 2: ")
	responsePart2 := ""
	for stackNumber, stack := range stacksPart2 {
		fmt.Println(stackNumber, " ", stack)
		responsePart2 += stack[len(stack)-1]
	}
	fmt.Println("\nResponse: " + responsePart2)

}
