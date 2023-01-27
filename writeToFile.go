package main

import (
	"os"
	"strconv"
	"strings"
)

func writeAllWordCountToFile(file os.DirEntry, min int, dir string, f os.File, c chan string) {
	fileLocation := dir + file.Name()
	// Convert slice of strings to string to bytes
	words := strings.Join(getWordCount(fileLocation, min), "")
	f.Write([]byte(file.Name() + "\n" + words + "\n"))
	c <- fileLocation + " counted!"
}

func writeSingleWordCountToFile(file os.DirEntry, searchWord string, dir string, f os.File, c chan string) {
	fileLocation := dir + file.Name()
	number := getSingleWordCount(fileLocation, searchWord)
	num := strconv.Itoa(number)
	f.Write([]byte(num + " times in " + file.Name() + "\n"))
	c <- searchWord + " appears " + num + " times in " + fileLocation
}
