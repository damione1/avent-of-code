package main

import (
	"fmt"
	"os"
)

func main() {
	d, err := os.ReadFile("puzzles/2022_12_06/dataset.txt")
	if err != nil {
		panic(err)
	}

	data := []rune(string(d))

	windowLength := 4 //set 14 for part 2
	pointValues := map[rune]int{}

	for _, c := range data[:windowLength] {
		pointValues[c]++
	}

	for i, c := range data[windowLength:] {
		pointValues[data[i]]--
		pointValues[c]++
		if checkAllUnique(pointValues) {
			fmt.Println(i + windowLength + 1)
			break
		}
	}

}

func checkAllUnique(m map[rune]int) bool {
	for _, v := range m {
		if v >= 2 {
			return false
		}
	}
	return true
}
