package main

import (
	"advent-of-code-21/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func A() {
	utils.SetBasePathToCurrentDir()

	// get the required file
	file, err := os.Open("input.txt")
	utils.HandleError(err, "opening input file")
	defer func(file *os.File) {
		err := file.Close()
		utils.HandleError(err, "closing input file")
	}(file)

	// stage it
	signals := make([]Signal, 0)
	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		currentLine := strings.TrimSpace(scanner.Text())
		signals = append(signals, NewSignal(currentLine))
	}

	// digit -> unique segments count
	// 1 -> 2
	// 4 -> 4
	// 7 -> 3
	// 8 -> 7
	uniqueSegmentCounts := []int{2, 4, 3, 7}
	countOfOutputWithUniqueSegmentDigit := 0
	for _, s := range signals {
		for _, o := range s.output {
			if utils.InList(len(o), uniqueSegmentCounts) {
				countOfOutputWithUniqueSegmentDigit++
			}
		}
	}

	fmt.Println("Count of output digits with unique number of segments: ", countOfOutputWithUniqueSegmentDigit)
}

func B() {
	utils.SetBasePathToCurrentDir()

	// get the required file
	file, err := os.Open("input.txt")
	utils.HandleError(err, "opening input file")
	defer func(file *os.File) {
		err := file.Close()
		utils.HandleError(err, "closing input file")
	}(file)

	// stage it
	signals := make([]Signal, 0)
	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		currentLine := strings.TrimSpace(scanner.Text())
		signals = append(signals, NewSignal(currentLine))
	}

	// segments count -> possible digits
	// 2 -> 1
	// 3 -> 7
	// 4 -> 4
	// 5 -> 2, 3, 5
	// 6 -> 0, 6, 9
	// 7 -> 8
	sum := 0

	// loop through all signals
	for _, s := range signals {
		digitToStr := make(map[string]string) // < digit, str_representation >

		// stage map with empty strings
		for digit := 0; digit < 10; digit++ {
			digitToStr[utils.Int2String(digit)] = ""
		}

		// closed function which returns whether all digits have been found
		allDigitsFound := func() bool {
			for _, v := range digitToStr {
				if v == "" {
					return false
				}
			}
			return true
		}

		for !allDigitsFound() {
			// loop through the ten unique signal patterns
			for _, o := range s.defs {
				// check which digit is being display based on the length of string representation o
				switch len(o) {
				case 2:
					// digit is 1
					digitToStr["1"] = o
				case 3:
					// digit is 7
					digitToStr["7"] = o
				case 4:
					// digit is 4
					digitToStr["4"] = o
				case 5:
					// digit could be 2,3,5
					/*
						to distinguish between these we need to check certain segment properties
						that are unique amongst string representations of 2,3,5:
							if (1 ⊂ o)
								then digit is 3
							else if len(2 U o) == 7 i.e activates all segments
								then digit is 2
							else
								digit is 5
					*/
					if len(digitToStr["1"]) > 0 && len(digitToStr["4"]) > 0 && len(digitToStr["8"]) > 0 { // we rely on already have found the mappings for 1,4,8
						if utils.IsSubset(o, digitToStr["1"]) {
							digitToStr["3"] = o
						} else if len(utils.StringUnion(o, digitToStr["4"])) == 7 {
							digitToStr["2"] = o
						} else {
							digitToStr["5"] = o
						}
					}
				case 6:
					// digit could be 0,6,9
					/*
						to distinguish between these we need to check certain segment properties
						that are unique amongst string representations of 0,6,9:
							if len(o U 7) == 7 i.e activates all segments
								then digit is 6
							else if len(o U 4) == 7 i.e activates all segments
								then digit is 0
							else
								digit is 9
					*/
					if len(digitToStr["7"]) > 0 && len(digitToStr["4"]) > 0 && len(digitToStr["8"]) > 0 { // we rely on already have found the mappings for 7,4,8
						if len(utils.StringUnion(o, digitToStr["7"])) == 7 {
							digitToStr["6"] = o
						} else if len(utils.StringUnion(o, digitToStr["4"])) == 7 {
							digitToStr["0"] = o
						} else {
							digitToStr["9"] = o
						}
					}
				case 7:
					// digit is 8
					digitToStr["8"] = o
				}
			}
		}

		// Now that all mappings have been found, translate the output value for this singal
		outputVal := ""
		for _, digit := range s.output {
			outputVal += findDigitFromStringRepr(digitToStr, digit)
		}
		// and add to global sum
		sum += utils.String2Int(outputVal)
	}

	fmt.Println("Sum of all output values: ", sum)
}

func findDigitFromStringRepr(digitToStr map[string]string, repr string) string {

	for k, v := range digitToStr {
		if utils.IsSubset(repr, v) && len(repr) == len(v) {
			return k
		}
	}
	log.Fatalf("Digit for str repr [%s] not found", repr)
	return ""
}
