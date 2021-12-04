package main

import (
	"advent-of-code-21/utils"
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type charCount struct{
	char string
	freq int
}

type ColCharCounts []*charCount

func (d *ColCharCounts) IncColCharCountFor(char string){
	found := false
	for _,cc := range *d {
		if cc.char == char {
			cc.freq++
			found=true
		}
	}

	if !found {
		*d = append(*d, &charCount{char, 1})
	}
}

// Len is part of sort.Interface.
func (d ColCharCounts) Len() int {
	return len(d)
}

// Swap is part of sort.Interface.
func (d ColCharCounts) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

// Less is part of sort.Interface. We use count as the value to sort by
func (d ColCharCounts) Less(i, j int) bool {
	return d[i].freq < d[j].freq
}

func A() {
	utils.SetBasePathToCurrentDir()

	file, err := os.Open("input.txt")

	utils.HandleError(err, "while opening input file")

	charCounts:=  make(map[int]ColCharCounts) // <col_no, <char, count> >

	scanner := bufio.NewScanner(file)
	numberOfCols := 0
	for row:=0; scanner.Scan();row++ {
		currentLine  := strings.TrimSpace(scanner.Text())

		for col, c := range currentLine {
			if charCounts[col] == nil{
				charCounts[col] = ColCharCounts{}
			}

			colCharCounts:=charCounts[col]

			colCharCounts.IncColCharCountFor(string(c))
			charCounts[col]=colCharCounts
			//fmt.Println()
		}

		numberOfCols = len(currentLine)
	}
	gamRate := ""
	epsRate := ""

	for i:=0; i<numberOfCols;i++ {
		cc := charCounts[i]
		sort.Sort(cc)

		//fmt.Println(cc)

		min, max := getMaxMinFreqCharForColumn(cc)

		gamRate = gamRate+max
		epsRate = epsRate+min
	}

	fmt.Println("Gamma rate: ", utils.Binary2Int(gamRate))
	fmt.Println("Epsilon rate: ", utils.Binary2Int(epsRate))
	fmt.Println("Ans: ",  utils.Binary2Int(gamRate)*utils.Binary2Int(epsRate))
}

func getMaxMinFreqCharForColumn(freqList ColCharCounts) (string, string){
	min :=  freqList[0].freq
	max :=  freqList[0].freq

	minChar := freqList[0].char
	maxChar := freqList[0].char
	for _,v := range freqList{
		if v.freq > max {
			max = v.freq
			maxChar = v.char
		}
		if v.freq < min {
			min = v.freq
			minChar = v.char
		}
	}

	return minChar, maxChar
}

