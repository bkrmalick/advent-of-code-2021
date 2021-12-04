package main

import (
	"advent-of-code-21/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func B() {
	utils.SetBasePathToCurrentDir()

	file, err := os.Open("input.txt")
	utils.HandleError(err, "opening input file")
	defer func(file *os.File) {
		err := file.Close()
		utils.HandleError(err, "closing input file")
	}(file)

	allLines := []string{}

	scanner := bufio.NewScanner(file)
	for row := 0; scanner.Scan(); row++ {
		currentLine := strings.TrimSpace(scanner.Text())
		allLines = append(allLines, currentLine)
	}

	mcv := func(chars []string, colNumber int) string {
		_, max := getMinMaxFreqChars(chars, colNumber, "1")
		return max
	}
	oxyRating := filterByCriteria(0, allLines, mcv)

	lcv := func(chars []string, colNumber int) string {
		min, _ := getMinMaxFreqChars(chars, colNumber, "0")
		return min
	}

	scrubRating := filterByCriteria(0, allLines, lcv)

	fmt.Println("oxygen generator rating: ", utils.Binary2Int(oxyRating))
	fmt.Println("CO2 scrubber rating: ", utils.Binary2Int(scrubRating))
	fmt.Println("Ans: ", utils.Binary2Int(oxyRating)*utils.Binary2Int(scrubRating))
}

// given a starting colIndex and inp of list of strings, recursively loops through each of the columns
// and filters out inp strings which do not match the string given by criteria function
// returns last standing string from the input
func filterByCriteria(colIndex int, inp []string, criteria func(chars []string, colNumber int) string) string {
	if len(inp) == 1 {
		return inp[0]
	}

	// get characters in column number colIndex
	currentCol := []string{}
	for _, k := range inp {
		currentCol = append(currentCol, string(k[colIndex]))
	}

	// get the required answer for this column
	criteriaAnsCharacter := criteria(inp, colIndex)

	// filter out those strings from inp which do not match the
	// answer
	newInp := []string{}
	for _, v := range inp {
		if string(v[colIndex]) == criteriaAnsCharacter {
			newInp = append(newInp, v)
		}
	}

	// do the remaining columns
	return filterByCriteria(colIndex+1, newInp, criteria)
}
