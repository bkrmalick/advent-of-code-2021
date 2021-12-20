package main

import (
	"advent-of-code-21/utils"
	"bufio"
	"fmt"
	"github.com/montanaflynn/stats"
	"log"
	"os"
	"reflect"
	"strings"
)

func AB() {
	utils.SetBasePathToCurrentDir()

	// get the required file
	file, err := os.Open("input.txt")
	utils.HandleError(err, "opening input file")
	defer func(file *os.File) {
		err := file.Close()
		utils.HandleError(err, "closing input file")
	}(file)

	// stage it
	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		currentLine := strings.TrimSpace(scanner.Text())
		lines = append(lines, currentLine)
	}

	// declare all possible brackets, and extract opening/closing characters into separate arrays
	bracketPairToCorruptionScore := map[string]int{
		"{}": 1197,
		"[]": 57,
		"()": 3,
		"<>": 25137,
	}
	bracketToCompletionScore := map[string]int{
		")": 1,
		"]": 2,
		"}": 3,
		">": 4,
	}
	allPossibleBracketPairs := make([]string, 0)
	for _, k := range reflect.ValueOf(bracketPairToCorruptionScore).MapKeys() {
		allPossibleBracketPairs = append(allPossibleBracketPairs, k.String())
	}
	var opening []string
	var closing []string
	var closingToOpeningBracketMap = make(map[string]string)
	for _, bracket := range allPossibleBracketPairs {
		openingBracket := string(bracket[0])
		closingBracket := string(bracket[1])
		opening = append(opening, openingBracket)
		closing = append(closing, closingBracket)
		closingToOpeningBracketMap[closingBracket] = openingBracket
	}

	offendingChars := make([]string, 0)
	completionScores := make([]int, 0)

	for _, l := range lines {
		isCorruptedLine := false
		for charIndex, c := range l {
			char := string(c)

			if !utils.InList(char, opening) && !utils.InList(char, closing) {
				log.Fatalf("Found a unexpected character [%s]", char)
			}

			if utils.InList(char, closing) {
				// this char is a closing bracket, validate the chunk enclosed by it
				matchingOpening := closingToOpeningBracketMap[char]
				valid := validateChunk(matchingOpening, charIndex-1, l, allPossibleBracketPairs)

				if !valid {
					fmt.Printf("Found corrupted char %s when closing at index %s \n", char, utils.Int2String(charIndex))
					offendingChars = append(offendingChars, char)
					isCorruptedLine = !valid
					break
				}
			}
		}

		if !isCorruptedLine {
			// line is not corrupted but incomplete
			currentLineCompletionScore := 0
			completingBracketsNeeded := findCompletingBrackets(l, allPossibleBracketPairs)

			for _, c := range completingBracketsNeeded {
				currentLineCompletionScore *= 5
				currentLineCompletionScore += bracketToCompletionScore[c]
			}

			completionScores = append(completionScores, currentLineCompletionScore)
		}
	}

	syntaxErrorScore := calculateCorruptionScore(offendingChars, bracketPairToCorruptionScore)
	fmt.Println("Total Syntax Error Score: ", syntaxErrorScore)

	overallCompletionScore, _ := stats.Median(stats.LoadRawData(completionScores))
	fmt.Printf("Overall Completion Score: %v", int(overallCompletionScore))

}

func calculateCorruptionScore(chars []string, scores map[string]int) int {
	totalScore := 0
	for _, c := range chars {

		for bracketPair, score := range scores {
			if strings.Contains(bracketPair, c) {
				totalScore += score
				break
			}
		}

	}
	return totalScore
}

// given a line and a stoppingChar (i.e. opening bracket that starts the chunk)
// find out if it is valid, if not return the offending bracket that was never closed within this chunk
func validateChunk(stoppingChar string, startingIndex int, fullLine string, allPossibleBracketPairs []string) bool {
	/*
		prepare a map that maintains counters of each bracket pair of the form

		map[string]int{
			"<>" :0,
			"{}" :0,
			"()" :0,
			"[]" :0,
		}
	*/
	counters := map[string]int{}
	for _, bracketPair := range allPossibleBracketPairs {
		counters[bracketPair] = 0
	}

	done := false

	// update the count map for each type of pair
	// up until the stoppingChar (opening bracket for chunk) is found
	for i := startingIndex; i > 0 && !done; i-- {
		currentChar := string(fullLine[i])
		for k, v := range counters {
			if strings.Contains(k, currentChar) {
				ind := strings.Index(k, currentChar)
				if ind == 0 {
					// opening bracket
					if currentChar == stoppingChar && v == 0 {
						done = true
						break
					}
					counters[k]++
				} else {
					// closing bracket
					counters[k]--
				}
			}
		}
	}

	// if any type of pair had a non-zero count for this chunk, it is invalid
	for _, count := range counters {
		if count != 0 {
			return false
		}
	}

	return true
}

// given a line, iterate through it in reverse (right to left) and find brackets which weren't closed,
// returns a string arrray that will close them in the correct order
func findCompletingBrackets(fullLine string, allPossibleBracketPairs []string) []string {
	/*
		prepare a map that maintains counters of each bracket pair of the form

		map[string]int{
			"<>" :0,
			"{}" :0,
			"()" :0,
			"[]" :0,
		}
	*/
	counters := map[string]int{}
	for _, bracketPair := range allPossibleBracketPairs {
		counters[bracketPair] = 0
	}
	completingBrackets := make([]string, 0)
	// iterate over line in reverse making sure to update the count map for each type of pair
	for i := len(fullLine) - 1; i >= 0; i-- {
		currentChar := string(fullLine[i])
		for k, v := range counters {
			if strings.Contains(k, currentChar) {
				ind := strings.Index(k, currentChar)
				if ind == 0 {
					// opening bracket
					if v == 0 {
						// if value of the counter for this type of bracket is zero,
						// then we know it was never closed because we're iterating the line right to left
						completingBrackets = append(completingBrackets, string(k[1]))
					} else {
						counters[k]++
					}
				} else {
					// closing bracket
					counters[k]--
				}
			}
		}
	}
	return completingBrackets
}
