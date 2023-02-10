package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"golang.org/x/exp/slices"
)

func getNames(dir string) {
	f, err := os.Create(dir + "/names.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	content, err := ioutil.ReadFile("./names.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	lowerCaseNames := strings.ToLower(string(content))
	names := strings.Split(lowerCaseNames, "\n")

	// Create an array of all the word maps that can then be merged into a single map and written to the file
	bigWordMap := []map[string]int{}
	readFilesWithoutNewFile(dir, *f, &bigWordMap)
	mergedMap := mergeMaps(bigWordMap)

	// Create a new map that only contains names
	namesMap := make(map[string]int)


	// Remove all the keys from the merged map that aren't in the names array
	for key, element := range mergedMap {
		if slices.Contains(names, key) {
			fmt.Println(true)
			namesMap[key] = element
		}
	}

	writeNameCountToFile(namesMap, *f)
	defer f.Close()

}

