
	package main

	import (
		"bufio"
		"log"
		"os"
	)

	func main() {
		content, err := os.Open("puzzles/2022_12_09/dataset.txt")

		if err != nil {
			log.Fatal(err)
		}

		scanner := bufio.NewScanner(content)

		var lines []string
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}

		//write your code here
	}