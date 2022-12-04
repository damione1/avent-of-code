package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func main() {

	nextDate, err := get_next_available_date()

	if err != nil {
		log.Fatal(err)
	}

	err = create_puzzle_file(nextDate)
	if err != nil {
		log.Fatal(err)
	}

	err = get_puzzle(nextDate)
	if err != nil {
		log.Fatal(err)
	}

}

func getSession() string {
	return os.Getenv("AOC_SESSION")
}

func get_puzzle(input time.Time) error {

	url := "https://adventofcode.com/" + input.Format("2006") + "/day/" + input.Format("2") + "/input"

	spaceClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Cookie", "session="+getSession())
	req.Header.Set("User-Agent", "AdventOfCode")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		return getErr
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return readErr
	}

	err = os.WriteFile("puzzles/"+input.Format("2006_01_02")+"/dataset.txt", body, 0644)
	if err != nil {
		return err
	}

	fmt.Println("Downloaded puzzle for " + input.Format("2006-01-02"))
	return nil
}

func create_puzzle_file(input time.Time) error {
	err := os.MkdirAll("puzzles/"+input.Format("2006_01_02"), os.ModePerm)
	if err != nil {
		return err
	}
	goFile, err := os.Create("puzzles/" + input.Format("2006_01_02") + "/main.go")
	if err != nil {
		return err
	}

	_, err = goFile.WriteString(`
	package main

	import (
		"bufio"
		"log"
		"os"
	)

	func main() {
		content, err := os.Open("puzzles/` + input.Format("2006_01_02") + `/dataset.txt")

		if err != nil {
			log.Fatal(err)
		}

		scanner := bufio.NewScanner(content)

		var lines []string
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}

		//write your code here
	}`)
	if err != nil {
		return err
	}

	fmt.Println("Created folder for " + input.Format("2006-01-02"))

	return nil
}

func get_next_available_date() (time.Time, error) {
	files := []string{}
	filepath.Walk("puzzles", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			files = append(files, info.Name())
		}
		return nil
	})

	lastFolder := files[len(files)-1]
	lastDate, _ := time.Parse("2006_01_02", lastFolder)

	timeNow := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, lastDate.Location())
	nextDate := time.Date(lastDate.Year(), lastDate.Month(), lastDate.Day(), 0, 0, 0, 0, lastDate.Location()).Add(time.Hour * 24)

	if nextDate.Year() != timeNow.Year() {
		nextDate = timeNow
	}

	if nextDate.After(time.Date(nextDate.Year(), 12, 25, 0, 0, 0, 0, nextDate.Location())) {
		return time.Time{}, fmt.Errorf("nextDate is after 25th of December. Everything is done")
	} else if nextDate.Before(time.Date(nextDate.Year(), 12, 1, 0, 0, 0, 0, nextDate.Location())) {
		return time.Time{}, fmt.Errorf("The puzzles for this year is not available yet")
	} else if nextDate.After(timeNow) {
		return time.Time{}, fmt.Errorf("No new puzzle available yet. Come back on %s", nextDate.Format("2006-01-02"))
	}

	fmt.Println("Next available puzzle is " + nextDate.Format("2006-01-02"))
	return nextDate, nil
}
