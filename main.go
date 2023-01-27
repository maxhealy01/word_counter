package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Read all the file names in a directory
	dir := os.Args[1] + "/"
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Supply a directory for the first argument")
		return
	}
	// Create a file within the directory to write the count to
	f, err := os.Create(dir + "word_count.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer f.Close()

	// Convert user-supplied number to int
	// this will be the minimum number of repetitions
	// for a word to be recorded in the word count.
	min, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Supply a number for the second argument")
		return
	}
	for _, file := range files {
		fileLocation := dir + file.Name()
		f.WriteString(file.Name() + "\n")
		// Convert slice of strings to string to bytes
		words := strings.Join(getWordCount(fileLocation, min), "")
		f.Write([]byte(words))
	}
}
