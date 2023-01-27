package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func getWordCount(s string, min int) []string {
	f, err := os.Open(s)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	defer f.Close()

	// Perform necessary conversions to string
	// if ODT, doc, etc
	var text string
	switch {
	case strings.HasSuffix(f.Name(), ".odt"):
		text = odtToString(f)
	case strings.HasSuffix(f.Name(), ".doc"):
		text = docToString(f)
	case strings.HasSuffix(f.Name(), ".docx"):
		text = docxToString(f)
	case strings.HasSuffix(f.Name(), ".pdf"):
		text = pdfToString(f)
	}
	words := stringToSlice(text)
	return mapWordsAndSort(words, min)
}

func stringToSlice(text string) []string {
	// Remove special characters
	text = regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(text, "")
	text = strings.ToLower(text)
	return strings.Split(text, " ")
}

// Create a slice of strings in ascending value
// starting at a minimum number of word repetitions.
func mapWordsAndSort(words []string, min int) []string {
	wordsMap := map[string]int{}

	for _, word := range words {
		wordsMap[word]++
	}

	delete(wordsMap, "")
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
		if wordsMap[k] >= min {
			value := strconv.Itoa(wordsMap[k])
			s := value + ": " + k + "\n"
			wordCounts = append(wordCounts, s)
		}
	}
	return wordCounts
}
