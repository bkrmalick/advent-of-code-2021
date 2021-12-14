package main

import (
	"advent-of-code-21/utils"
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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
	heightMap := make([][]int, 0)
	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		currentLine := strings.TrimSpace(scanner.Text())
		strLetters := strings.Split(currentLine, "")
		heightMap = append(heightMap, make([]int, len(currentLine)))
		for j, letter := range strLetters {
			heightMap[i][j] = utils.String2Int(letter)
		}
	}

	sumRiskLevels := 0
	basinSizes := make([]int, 3)
	alreadyCheckedPoints := make([]Point, 0)

	for i, row := range heightMap {
		for j := range row {
			point := Point{i, j}
			if isLowPoint(heightMap, point) {
				sumRiskLevels += heightMap[i][j] + 1
				size := getBasinSizeFromLowPoint(
					heightMap,
					point,
					heightMap[i][j]-1, // bit hacky to always make sure it always runs for first iteration
					&alreadyCheckedPoints,
				)
				basinSizes = append(basinSizes, size)
			}
		}
	}

	fmt.Println("part a: ", sumRiskLevels)

	sort.Slice(basinSizes, func(i, j int) bool {
		return basinSizes[i] > basinSizes[j]
	})

	fmt.Println("part b: ", basinSizes[0]*basinSizes[1]*basinSizes[2])

}

func getBasinSizeFromLowPoint(heightMap [][]int, p Point, prevValue int, alreadyChecked *[]Point) int {
	currentVal := getValue(heightMap, p)

	if !areValidCoordinates(heightMap, p) ||
		currentVal == 9 ||
		(currentVal < prevValue) ||
		isPointAlreadyChecked(*alreadyChecked, p) {
		return 0
	}

	*alreadyChecked = append(*alreadyChecked, p)

	// recursively call all possible paths from this point (they'll return 0 if they're not valid)
	up := getBasinSizeFromLowPoint(heightMap, Point{p.i - 1, p.j}, currentVal, alreadyChecked)
	left := getBasinSizeFromLowPoint(heightMap, Point{p.i, p.j - 1}, currentVal, alreadyChecked)
	down := getBasinSizeFromLowPoint(heightMap, Point{p.i + 1, p.j}, currentVal, alreadyChecked)
	right := getBasinSizeFromLowPoint(heightMap, Point{p.i, p.j + 1}, currentVal, alreadyChecked)

	return 1 + left + right + up + down
}

func isPointAlreadyChecked(alreadyChecked []Point, point Point) bool {
	for _, p := range alreadyChecked {
		if p.j == point.j && p.i == point.i {
			return true
		}
	}
	return false
}

func isLowPoint(heightMap [][]int, p Point) bool {
	adjacentValues := []int{
		getValue(heightMap, Point{p.i - 1, p.j}),
		getValue(heightMap, Point{p.i, p.j - 1}),
		getValue(heightMap, Point{p.i + 1, p.j}),
		getValue(heightMap, Point{p.i, p.j + 1}),
	}
	if getValue(heightMap, p) < utils.Min(adjacentValues...) {
		return true
	}
	return false
}

func areValidCoordinates(heightMap [][]int, p Point) bool {
	height := len(heightMap)
	width := len(heightMap[0])
	if p.j < 0 || p.j >= width || p.i < 0 || p.i >= height {
		return false
	}
	return true
}

func getValue(heightMap [][]int, p Point) int {
	if !areValidCoordinates(heightMap, p) {
		return math.MaxInt32
	}
	return heightMap[p.i][p.j]
}
