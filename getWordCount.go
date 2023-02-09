package main

import (
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func getWordCount(f string, min int) []string {
	text := openAndConvertFileToString(f)
	words := stringToSlice(text)
	wordsMap := mapWordsAndSort(words)
	return wordCountStrings(wordsMap, min)
}

func stringToSlice(text string) []string {
	// Remove special characters
	text = regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(text, "")
	text = strings.ToLower(text)
	return strings.Split(text, " ")
}

// Create a slice of strings in ascending value
// starting at a minimum number of word repetitions.
func mapWordsAndSort(words []string) map[string]int {
	wordsMap := map[string]int{}

	for _, word := range words {
		wordsMap[word]++
	}

	// Delete common words, whose repetition isn't very interesting!
	commonWords := []string{"", "to", "and", "i", "in", "of", "the", "that", "is", "be", "a", "it", "am", "or", "by", "with", "have", "this", "my", "for", "but", "so", "as"}

	for _, word := range commonWords {
		delete(wordsMap, word)
	}

	return wordsMap
}

func wordCountStrings(wordsMap map[string]int, min int) []string {
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
