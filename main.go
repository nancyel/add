package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Usage: ./add init|'phrase'")
		return
	}

	switch args[0] {
	case "init":
		createNewFile()
	default:
		addPhrase(args[0])
	}
}

func createNewFile() {
	dir := os.Getenv("HOME") + "/Desktop/phrases"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, os.ModePerm)
	}

	filename := time.Now().Format("20060102") + ".md"
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

func addPhrase(phrase string) {
	dir := os.Getenv("HOME") + "/Desktop/phrases"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		fmt.Println("Phrase directory does not exist. Create a new file with 'add init' command first.")
		return
	}

	filename := time.Now().Format("20060102") + ".md"
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
