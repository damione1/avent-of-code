package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
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

	var stackSchema []string
	for _, line := range lines {
		if line == "" {
			break
		}
		for strings.Contains(line, "]    ") {
			line = strings.ReplaceAll(line, "]    ", "] [_]")
		}
		for strings.Contains(line, "    [") {
			line = strings.ReplaceAll(line, "    [", "[_] [")
		}
		line = strings.ReplaceAll(line, "[", "")
		line = strings.ReplaceAll(line, "]", "")
		line = strings.ReplaceAll(line, " ", "")
		stackSchema = append(stackSchema, line)
	}

	stacksPart1 := make([][]string, len(stackSchema[len(stackSchema)-1]))
	stacksPart2 := make([][]string, len(stackSchema[len(stackSchema)-1]))
	for _, stackLine := range stackSchema {
		for i, stack := range stackLine {
			if stack == '_' {
				continue
			}
			if _, err := strconv.Atoi(string(stack)); err == nil {
				continue
			}
			stacksPart1[i] = append([]string{string(stack)}, stacksPart1[i]...)
			stacksPart2[i] = append([]string{string(stack)}, stacksPart2[i]...)
		}
	}

	//print stacks
	fmt.Println("Original stacks:")
	for stackNumber, stack := range stacksPart1 {
		fmt.Println(stackNumber, stack)
	}

	instructions := false
	for _, line := range lines {
		if line == "" {
			instructions = true
			continue
		}
		if !instructions {
			continue
		}
		regex := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
		matches := regex.FindStringSubmatch(line)
		if len(matches) != 4 {
			log.Fatal("Wrong line format: ", line)
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
