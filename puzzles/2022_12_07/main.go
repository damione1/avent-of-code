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

type directory struct {
	name     string
	children []*directory
	files    []file
	parent   *directory
	size     int
}

type file struct {
	name string
	size int
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

	var folderSizeMap []int
	directory := generateDirectoryTree(instructions, &folderSizeMap)

	printTree(*directory, 0)

	var usedSpace int = 0
	var folderSum int = 0
	for _, size := range folderSizeMap {
		if size <= 100000 {
			folderSum += size
		}
		if size > usedSpace {
			usedSpace = size
		}
	}
	fmt.Println("Sum of the directories with size less than 100000: ", usedSpace)

	var totalSpace int = 70000000
	var freeSpace int = totalSpace - usedSpace
	var requiredSpace int = 30000000 - freeSpace

	sort.Ints(folderSizeMap)

	fmt.Println("Required space: ", requiredSpace)
	fmt.Println("Free space: ", freeSpace)
	for _, size := range folderSizeMap {
		if size >= requiredSpace {
			fmt.Println("Smallest folder that can be deleted to free up the required space: ", size)
			break
		}
	}

}

func printTree(folder directory, level int) {
	fmt.Println(strings.Repeat("-", level), "\033[1m"+folder.name, "(", folder.size, ")", "\033[0m")
	for _, file := range folder.files {
		fmt.Println(strings.Repeat(" ", level), "| "+file.name, file.size)
	}
	for _, child := range folder.children {
		printTree(*child, level+1)
	}
}

func generateDirectoryTree(instructions []string, folderSizeMap *[]int) *directory {
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
					size: size,
				}
				currentDirectoryPointer.files = append(currentDirectoryPointer.files, file)
			}
		}

	}

	calcDirectoriesSize(&root, folderSizeMap)

	return &root

}

func calcDirectoriesSize(folder *directory, folderSizeMap *[]int) int {
	var sum int = 0
	for _, file := range folder.files {
		sum += file.size
	}
	for _, child := range folder.children {
		sum += calcDirectoriesSize(child, folderSizeMap)
	}
	folder.size = sum
	*folderSizeMap = append(*folderSizeMap, sum)
	return sum
}
