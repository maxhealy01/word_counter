package main

func getSingleWordCount(f string, sw string) int {
	text := openAndConvertFileToString(f)
	words := stringToSlice(text)
	return countSingleWord(words, sw)
}

func countSingleWord(words []string, sw string) int {
	count := 0
	for _, word := range words {
		if word == sw {
			count++
		}
	}
	return count
}
