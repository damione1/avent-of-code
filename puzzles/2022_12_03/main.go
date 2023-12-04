package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	dataset, _ := os.Open("puzzles/2022_12_03/dataset.txt")
	scanner := bufio.NewScanner(dataset)

	var rucksacks []string
	for scanner.Scan() {
		rucksacks = append(rucksacks, scanner.Text())
	}

	sumPart1 := 0
	for _, rucksack := range rucksacks {
		side1, side2 := rucksack[:len(rucksack)/2], rucksack[len(rucksack)/2:]
		for _, itemRune := range side1 {
			item := string(itemRune)
			if strings.Contains(side2, item) {
				sumPart1 += getPriority(item)
				break
			}
		}
	}
	fmt.Println("Result part 1: ", sumPart1)

	sumPart2 := 0
	groupSize := 3
	for i := 0; i < len(rucksacks)/groupSize; i++ {
		group := rucksacks[3*i : 3*i+groupSize]
		for _, itemRune := range group[0] {
			item := string(itemRune)
			if strings.Contains(group[1], item) && strings.Contains(group[2], item) {
				sumPart2 += getPriority(item)
				break
			}
		}
	}
	fmt.Println("Result part 2: ", sumPart2)
}

func getPriority(item string) int {
	priorities := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return strings.Index(priorities, item) + 1
}
