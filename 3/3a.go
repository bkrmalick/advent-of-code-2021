package main

import (
	"advent-of-code-21/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func A() {
	utils.SetBasePathToCurrentDir()

	file, err := os.Open("input.txt")
	utils.HandleError(err, "opening input file")
	defer func(file *os.File) {
		err := file.Close()
		utils.HandleError(err, "closing input file")
	}(file)

	allLines := []string{}

	scanner := bufio.NewScanner(file)
	numberOfCols := 0
	for row := 0; scanner.Scan(); row++ {
		currentLine := strings.TrimSpace(scanner.Text())
		allLines = append(allLines, currentLine)

		if len(currentLine) > numberOfCols {
			// assumption that all lines same number of columns
			// so this might be unnecessary
			numberOfCols = len(currentLine)
		}
	}
	gamRate := ""
	epsRate := ""

	for i := 0; i < numberOfCols; i++ {
		min, max := getMinMaxFreqChars(allLines, i, "0")
		gamRate = gamRate + max
		epsRate = epsRate + min
	}

	fmt.Println("Gamma rate: ", utils.Binary2Int(gamRate))
	fmt.Println("Epsilon rate: ", utils.Binary2Int(epsRate))
	fmt.Println("Ans: ", utils.Binary2Int(gamRate)*utils.Binary2Int(epsRate))
}

// given a list of binary numbers, finds the characters with the max and min frequency in a particular column
// assumes the characters are binary i.e. one of ["1", "0"] can provide a tiebreaker character to return when
// frequency of both characters are the same in the column
func getMinMaxFreqChars(ls []string, colNumber int, tiebreaker string) (string, string) {
	sum := 0

	for _, v := range ls {
		x, err := strconv.Atoi(string(v[colNumber]))
		utils.HandleError(err, "trying to parse string to an int")
		sum += x
	}

	numberOf1s := sum
	numberOf0s := len(ls) - sum

	if numberOf0s > numberOf1s {
		return "1", "0"
	} else if numberOf1s > numberOf0s {
		return "0", "1"
	} else {
		return tiebreaker, tiebreaker
	}
}
