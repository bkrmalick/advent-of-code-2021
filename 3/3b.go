package main

import (
	"advent-of-code-21/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func B() {
	utils.SetBasePathToCurrentDir()

	file, err := os.Open("input.txt")

	utils.HandleError(err, "while opening input file")

	charCounts:=  make(map[int]ColCharCounts) // <col_no, <char, count> >
	allLines := []string{}

	scanner := bufio.NewScanner(file)
	for row:=0; scanner.Scan();row++ {
		currentLine  := strings.TrimSpace(scanner.Text())
		allLines= append(allLines, currentLine)

		for col, c := range currentLine {
			if charCounts[col] == nil{
				charCounts[col] = ColCharCounts{}
			}

			colCharCounts:=charCounts[col]

			colCharCounts.IncColCharCountFor(string(c))
			charCounts[col]=colCharCounts
			//fmt.Println()
		}
	}

	oxyRatingData :=make([]string, len(allLines))
	copy(oxyRatingData,allLines)

	scrubRatingData := make([]string, len(allLines))
	copy(scrubRatingData,allLines)

	mcv:= func(chars []string, colNumber int) string {
		_, max := getMinMaxFreqChars(chars,colNumber, "1")
		return max
	}
	oxyRating := filterByCriteria(0,oxyRatingData, mcv)

	lcv:= func(chars []string, colNumber int) string {
		min, _ := getMinMaxFreqChars(chars,colNumber, "0")
		return min
	}

	scrubRating := filterByCriteria(0, scrubRatingData, lcv)

	fmt.Println("oxygen generator rating: ", utils.Binary2Int(oxyRating))
	fmt.Println("CO2 scrubber rating: ", utils.Binary2Int(scrubRating))
	fmt.Println("Ans: ",  utils.Binary2Int(oxyRating)*utils.Binary2Int(scrubRating))
}

func filterByCriteria(colIndex int, inp []string, criteria func(chars []string, colNumber int) string) string  {

	if len(inp) == 1{
		return inp[0]
	}
	firstCol := []string{}
	for _,k := range inp{
		firstCol = append(firstCol, string(k[colIndex]))
	}
	critAnsDigit := criteria(inp, colIndex)

	newInp := []string{}

	for _,v := range inp {
		if string(v[colIndex]) == critAnsDigit{
			newInp = append(newInp, v)
		}
	}

	return filterByCriteria(colIndex+1, newInp, criteria)
}

func getMinMaxFreqChars(ls []string, colNumber int,  tiebreaker string) (string, string){
	sum := 0

	for _,v := range ls{
		x,err :=strconv.Atoi(string(v[colNumber]))
		utils.HandleError(err, "trying to parse string to an int")
		sum+=x
	}

	numberOf1s := sum
	numberOf0s := len(ls) - sum

	if numberOf0s > numberOf1s {
		return "1", "0"
	} else 	if numberOf1s > numberOf0s {
		return "0","1"
	} else {
		return tiebreaker, tiebreaker
	}
}