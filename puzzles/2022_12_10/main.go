package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := os.Open("puzzles/2022_12_10/dataset.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(content)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	/* dataset
	   noop
	   addx -22
	   noop
	   noop
	   noop
	   noop
	   addx 3
	   addx 5
	   addx 2
	   addx -11
	   addx 16
	   addx -2
	*/
	cycleCount := 0
	xValue := 1
	signalSum := 0
	crt := ""
	for _, line := range lines {
		instructions := strings.Split(line, " ")

		signalSum = printX(xValue, cycleCount, signalSum)
		if instructions[0] == "noop" {
			cycleCount++
			crt = pixelPrint(cycleCount, xValue, crt)
			continue
		} else {
			cycleCount++
			crt = pixelPrint(cycleCount, xValue, crt)

			cycleCount++
			crt = pixelPrint(cycleCount, xValue, crt)
			signalSum = printX(xValue, cycleCount, signalSum)
			value, _ := strconv.Atoi(instructions[1])
			xValue += value
			continue
		}

	}

	fmt.Println("Signal sum: ", signalSum)

	for i := 0; i < len(crt); i++ {
		fmt.Print(string(crt[i]))
		if (i+1)%40 == 0 {
			fmt.Println("")
		}
	}

}

func printX(xValue int, cycleCount int, signalSum int) int {
	signalStrength := xValue * cycleCount
	if cycleCount == 20 || cycleCount == 60 || cycleCount == 100 || cycleCount == 140 || cycleCount == 180 || cycleCount == 220 {
		fmt.Println("Cycle count: ", cycleCount, " Signal strength: ", signalStrength)
		signalSum += signalStrength
	}

	return signalSum

}

func pixelPrint(cycle int, x int, crt string) string {

	pixelPosition := (cycle - 1) % 40
	spriteCenter := ((x - 1) % 40) + 1

	if pixelPosition >= spriteCenter-1 && pixelPosition <= spriteCenter+1 {
		crt += "#"
	} else {
		crt += "."
	}

	return crt

}
