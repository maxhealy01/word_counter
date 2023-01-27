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

	// Create a variable to hold the word
	// if the user supplies a word.
	var searchWord string
	// Convert user-supplied number to int
	// this will be the minimum number of repetitions
	// for a word to be recorded in the word count.
	min, err := strconv.Atoi(os.Args[2])
	if err != nil {
		searchWord = os.Args[2]
	}
	// Create a file within the directory to write the count to
	var newFileName string
	// Give the file a different name depending on what's being counted
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

	if min > 0 {
	for _, file := range files {
		go writeAllWordCountToFile(file, min, dir, *f, c)
	}} else {
		for _, file := range files {
			go writeSingleWordCountToFile(file, searchWord, dir, *f, c)
		}
	}

	for m := range files {
		fmt.Println(m, <- c)
	}
}

func writeAllWordCountToFile(file os.DirEntry, min int, dir string, f os.File, c chan string){
		fileLocation := dir + file.Name()
		// Convert slice of strings to string to bytes
		words := strings.Join(getWordCount(fileLocation, min), "")
		f.Write([]byte(file.Name() + "\n" + words + "\n"))
		c <- fileLocation + " counted!"
}

func writeSingleWordCountToFile(file os.DirEntry, searchWord string, dir string, f os.File, c chan string){
	fileLocation := dir + file.Name()
	number := getSingleWordCount(fileLocation, searchWord)
	num := strconv.Itoa(number)
	f.Write([]byte(num + " times in " + file.Name() + "\n"))
	c <- searchWord + " appears " + num + " times in " + fileLocation 
}