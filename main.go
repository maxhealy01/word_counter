package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	dir := os.Args[1]

	readFilesInDir(dir)
}

func readFilesInDir(dir string) {
	dir += "/"
	// Read all the file names in a directory
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Supply a directory for the first argument")
		return
	}

	// Create a variable to hold the word if the user supplies a word.
	var searchWord string
	// Convert user-supplied number to int. This will be the minimum number of repetitions
	// of repetitions for a word to be recorded in the word count.
	min, err := strconv.Atoi(os.Args[2])
	if err != nil {
		searchWord = os.Args[2]
	}
	// Create a file within the directory to write the count to.
	// Give the file a different name depending on what's being counted
	var newFileName string
	if min > 0 {
		newFileName = "word_count.txt"
	} else {
		newFileName = searchWord + ".txt"
	}

	// Add a channel so each file can have its own go routine
	c := make(chan string)
	for _, file := range files {
		if file.Type().IsDir() {
			fmt.Println(file.Type())
			readFilesInDir(dir + file.Name())
		} else {
			f, err := os.Create(dir + newFileName)
			if err != nil {
				fmt.Println("Hello: ", err)
				return
			}
			defer f.Close()
			if min > 0 {
				go writeAllWordCountToFile(file, min, dir, *f, c)

			} else {
				go writeSingleWordCountToFile(file, searchWord, dir, *f, c)

			}
		}
	}

	for m := range files {
		fmt.Println(m, <-c)
	}
}
