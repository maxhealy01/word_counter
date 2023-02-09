package main

import (
	"os"
	"sort"
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

func writeTotalCountToFile(bwm []map[string]int, f os.File) {
	mergedMaps := mergeMaps(bwm)
	wordCountStrings := strings.Join(wordCountStringsNoMin(mergedMaps), "")
	f.Write([]byte(wordCountStrings))
}

func wordCountStringsNoMin(wordsMap map[string]int) []string {
	// Sort map by value
	keys := make([]string, 0, len(wordsMap))

	for key := range wordsMap {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return wordsMap[keys[i]] < wordsMap[keys[j]]
	})

	wordCounts := []string{}
	for _, k := range keys {
		value := strconv.Itoa(wordsMap[k])
		s := value + ": " + k + "\n"
		wordCounts = append(wordCounts, s)
	}
	return wordCounts
}
