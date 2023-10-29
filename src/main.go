package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func searchInFile(keyword string, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", filename, err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 1

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, keyword) {
			fmt.Printf("Found '%s' in %s (line %d): %s\n", keyword, filename, lineNumber, line)
		}
		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file %s: %v\n", filename, err)
	}
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: search-word-in-files <word> <file1> [<file2> ...]")
		return
	}

	keyword := os.Args[1]

	for _, filename := range os.Args[2:] {
		searchInFile(keyword, filename)
	}
}

