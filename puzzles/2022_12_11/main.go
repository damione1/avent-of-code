package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	content, err := os.Open("puzzles/2022_12_11/dataset.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(content)

	//split rounds with empty line as separator
	var rounds []string
	var round string
	for scanner.Scan() {
		if scanner.Text() == "" {
			rounds = append(rounds, round)
			round = ""
			continue
		}
		round += scanner.Text()
	}

	for _, round := range rounds {

		//this is the string: Monkey 0:  Starting items: 57, 58  Operation: new = old * 19  Test: divisible by 7    If true: throw to monkey 2    If false: throw to monkey 3
		//keep only the numbers and the operation by replacing everything else with a space
		fmt.Sscanf(strings.NewReplacer(", ", ",", "* old", "^ 2").Replace(s),
			`Monkey %d:
  Starting items: %s
  Operation: new = old %s %d
  Test: divisible by %d
    If true: throw to monkey %d
    If false: throw to monkey %d`,
			&i, &items, &op, &v, &test, &t, &f)



}
