package main

import (
	"advent-of-code-21/utils"
	"bufio"
	"fmt"
	"os"
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

	// stage the energy levels in a 2D int array
	eLevels := make([][]int, 0)
	totalNumberOfLevels := 0
	for _, l := range lines {
		eLevelsRow := make([]int, 0)
		for _, c := range l {
			char := string(c)
			eLevelsRow = append(eLevelsRow, utils.String2Int(char))
			totalNumberOfLevels++
		}
		eLevels = append(eLevels, eLevelsRow)
	}

	const stepCount int = 100
	totalFlashCountForStepCount := 0
	syncFound := false

	printLevels(eLevels)
	for i := 0; !syncFound; i++ {
		flashCount := step(eLevels)

		if i < stepCount {
			totalFlashCountForStepCount += flashCount
		}

		fmt.Printf("Step %s: \n", utils.Int2String(i+1))

		if flashCount == totalNumberOfLevels {
			fmt.Println("SYNCED")
			syncFound = true
		}

		printLevels(eLevels)
	}

	fmt.Println("TOTAL FLASH COUNT: ", totalFlashCountForStepCount)
}

func printLevels(levels [][]int) {
	for _, row := range levels {
		for _, v := range row {
			fmt.Printf("%s,", utils.Int2String(v))
		}
		fmt.Println()
	}
	fmt.Println()
}

func step(levels [][]int) int {
	increaseLevelsBy(levels, 1)
	flashCount := handleFlashes(levels)
	return flashCount
}

func increaseLevelsBy(levels [][]int, increment int) {
	for i, row := range levels {
		for j := range row {
			levels[i][j] += increment
		}
	}
}

func handleFlashes(levels [][]int) int {
	flashed := true
	flashCount := 0

	for flashed { // keep running until no more flashes
		flashed = false
		for i, row := range levels {
			for j := range row {
				if levels[i][j] > 9 {
					// flash
					flashed = true
					flashCount++

					// increment all 8 adjacent octopuses
					increaseLevelAt(levels, i-1, j)
					increaseLevelAt(levels, i, j-1)
					increaseLevelAt(levels, i-1, j-1)
					increaseLevelAt(levels, i+1, j-1)
					increaseLevelAt(levels, i-1, j+1)
					increaseLevelAt(levels, i+1, j)
					increaseLevelAt(levels, i, j+1)
					increaseLevelAt(levels, i+1, j+1)

					levels[i][j] = 0
				}
			}
		}
	}
	return flashCount
}

func increaseLevelAt(levels [][]int, i int, j int) {
	if areValidCoordinates(levels, i, j) && levels[i][j] != 0 { // flash only if this octopus hasn't already flashed
		levels[i][j]++
	}
}

func areValidCoordinates(levels [][]int, i int, j int) bool {
	height := len(levels)
	width := len(levels[0])
	if j < 0 || j >= width || i < 0 || i >= height {
		return false
	}
	return true
}
