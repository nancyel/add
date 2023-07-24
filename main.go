package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Usage: ./add init|'phrase'|grab")
		return
	}

	switch args[0] {
	case "init":
		createNewFile()
	case "grab":
		grab()
	default:
		addPhrase(args[0])
	}
}

func createNewFile() {
	dir := os.Getenv("HOME") + "/Desktop/phrases"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, os.ModePerm)
	}

	filename := time.Now().Format("20060102") + ".txt"
	filepath := dir + "/" + filename

	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		_, err := os.Create(filepath)
		if err != nil {
			fmt.Println("Failed to create file:", err)
			return
		}
		fmt.Println("Created new file:", filename)
	} else {
		fmt.Println("File already exists:", filename)
	}
}

func grab() {
	// Call the GrabPhrases function to read today's phrases
	phrases, err := GrabPhrases()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the phrases
	fmt.Println(phrases)
}

func GrabPhrases() ([]string, error) {
	// Get today's date in the format "20230330"
	dateStr := time.Now().Format("20060102")

	// Construct the filename of the phrases file
	filename := filepath.Join("phrases", dateStr+".txt")

	// Read the contents of the phrases file
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Split the contents of the file into lines
	lines := strings.Split(string(bytes), "\n")

	// Extract the phrases from the lines
	var phrases []string
	for _, line := range lines {
		// Ignore empty lines and lines that start with a timestamp
		if line == "" {
			continue
		}
		// check if line starts with a timestamp in the format "15:40:33"
		match, _ := regexp.MatchString(`^\d{2}:\d{2}:\d{2}\s+`, line)
		if match {
			// extract the rest of the line after the timestamp
			line = strings.TrimSpace(line[9:])

			// do further processing per line
		}

		// Add the phrase to the list
		phrases = append(phrases, line)
	}

	return phrases, nil
}

func addPhrase(phrase string) {
	dir := os.Getenv("HOME") + "/Desktop/phrases"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		fmt.Println("Phrase directory does not exist. Create a new file with './add init' command first.")
		return
	}

	filename := time.Now().Format("20060102") + ".txt"
	filepath := dir + "/" + filename

	f, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Failed to open file:", err)
		return
	}

	defer f.Close()

	t := time.Now().Format("15:04:33")
	if _, err := f.WriteString("\n" + t + " " + phrase); err != nil {
		fmt.Println("Failed to write to file:", err)
		return
	}

	fmt.Println("Added phrase to file:", filename)
}
