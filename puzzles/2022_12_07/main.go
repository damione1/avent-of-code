package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type directory struct {
	name     string
	children []*directory
	files    []file
	parent   *directory
	size     int64
}

type file struct {
	name string
	size int64
}

func main() {
	content, err := os.Open("puzzles/2022_12_07/dataset.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(content)

	var instructions []string
	for scanner.Scan() {
		instructions = append(instructions, scanner.Text())
	}

	directory := generateDirectoryTree(instructions)

	printFolder(*directory, 0)

	//calculate sum of directory of less than 100000
	var folderSize = make(map[string]int64)
	calcBigDirectorySize(directory, folderSize)
	//calc the sum of the directories
	var sum int64 = 0
	for _, size := range folderSize {
		//fmt.Println(folderName, size)
		if size < 100000 {
			sum += size
		}
	}
	fmt.Println("Sum of the directories with size less than 100000: ", sum)

}

func printFolder(folder directory, level int) {
	fmt.Println(strings.Repeat("-", level), "\033[1m"+folder.name, folder.size, "\033[0m")
	for _, file := range folder.files {
		fmt.Println(strings.Repeat(" ", level), "| "+file.name, file.size)
	}
	for _, child := range folder.children {
		printFolder(*child, level+1)
	}
}

func calcBigDirectorySize(folder *directory, folderSize map[string]int64) {
	folderSize[folder.name] = int64(folder.size)
	for _, child := range folder.children {
		calcBigDirectorySize(child, folderSize)
	}
}
func generateDirectoryTree(instructions []string) *directory {
	var root directory = directory{
		name:     "root",
		size:     0,
		parent:   nil,
		files:    []file{},
		children: []*directory{},
	}
	var currentDirectoryPointer = &root

	for _, instruction := range instructions {
		if instruction == "" {
			continue
		}
		splittedInstructions := strings.Split(instruction, " ")
		if splittedInstructions[0] == string('$') {
			if splittedInstructions[1] == string("cd") {
				if splittedInstructions[2] == string("..") {
					if currentDirectoryPointer.parent != nil {
						currentDirectoryPointer = currentDirectoryPointer.parent
					} else {
						fmt.Println("Error: cannot go back from root")
					}
				} else if splittedInstructions[2] == string("/") {
					currentDirectoryPointer = &root
				} else {
					for _, child := range currentDirectoryPointer.children {
						if child.name == splittedInstructions[2] {
							currentDirectoryPointer = child
							break
						}
					}
				}
			} else if splittedInstructions[1] == string("ls") {
				//fmt.Println("Listing", currentDirectoryPointer.name)
			}

		} else if splittedInstructions[0] == string("dir") {
			var dirExists bool = false
			for _, child := range currentDirectoryPointer.children {
				if child.name == splittedInstructions[1] {
					dirExists = true
					break
				}
			}
			if !dirExists {
				newDir := directory{
					name:     splittedInstructions[1],
					size:     0,
					parent:   currentDirectoryPointer,
					files:    []file{},
					children: []*directory{},
				}
				currentDirectoryPointer.children = append(currentDirectoryPointer.children, &newDir)
			}
		} else {
			var fileExists bool = false
			for _, file := range currentDirectoryPointer.files {
				if file.name == splittedInstructions[1] {
					fileExists = true
					break
				}
			}
			if !fileExists {
				size, _ := strconv.Atoi(splittedInstructions[0])
				file := file{
					name: splittedInstructions[1],
					size: int64(size),
				}
				currentDirectoryPointer.files = append(currentDirectoryPointer.files, file)
				currentDirectoryPointer.size += file.size
				calcDirectorySize(*currentDirectoryPointer)

			}
		}

	}

	return &root

}

func calcDirectorySize(directory directory) {
	//add the size of the directory itself to his parent.
	currentDirectoryPointer := &directory
	for currentDirectoryPointer.parent != nil {
		currentDirectoryPointer.parent.size += currentDirectoryPointer.size
		currentDirectoryPointer = currentDirectoryPointer.parent
	}
}
