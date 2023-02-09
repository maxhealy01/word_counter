package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	dir := os.Args[1]

	// If the user wants the total of the words over all files, create a different path
	// in order to avoid creating multiple files.
	if os.Args[2] == "total" {
		f, err := os.Create(dir + "/total.txt")
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		// Create an array of all the word maps that can then be merged into a single map and written to the file
		var bigWordMap = []map[string]int{}
		readFilesWithoutNewFile(dir, *f, &bigWordMap)
		writeTotalCountToFile(bigWordMap, *f)
		defer f.Close()

	} else {
		readFilesInDir(dir)
	}
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

	f, err := os.Create(dir + newFileName)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer f.Close()

	// Add a channel so each file can have its own go routine
	c := make(chan string)
	for _, file := range files {
		if file.Type().IsDir() {
			fmt.Println(file.Type())
			go readFilesInDir(dir + file.Name())
		} else {
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
