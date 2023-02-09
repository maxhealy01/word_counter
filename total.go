package main

import (
	"fmt"
	"os"
)

func readFilesWithoutNewFile(dir string, f os.File, bigWordMap *[]map[string]int) {
	dir += "/"
	// Read all the file names in a directory
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Supply a directory for the first argument")
		return
	}

	for _, file := range files {
		if file.Type().IsDir() {
			readFilesWithoutNewFile(dir+file.Name(), f, bigWordMap)
		} else {
			fileLocation := dir + file.Name()
			// Get the map of the words, and then add this to the big word map
			wordsMap := getWordMap(fileLocation)
			*bigWordMap = append(*bigWordMap, wordsMap)
		}
	}
}

func getWordMap(f string) map[string]int {
	text := openAndConvertFileToString(f)
	words := stringToSlice(text)
	wordsMap := mapWordsAndSort(words)
	return wordsMap
}

// This function will ultimately merge all the word maps into one big map
func mergeMaps[M ~map[K]V, K comparable, V int](src []M) M {
	merged := make(M)
	for _, m := range src {
		for k, v := range m {
			merged[k] += v
		}
	}
	return merged
}
