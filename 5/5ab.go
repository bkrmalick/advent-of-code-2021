package main

import (
	"advent-of-code-21/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type LinearFormula struct {
	// y = mx + c
	// y = 3
	m, c int

	// x = 3
	isVertical bool
	xvalue     int

	minX, maxX, minY, maxY int
}

func (f *LinearFormula) isPointOnLine(x int, y int) bool {
	if x < f.minX || x > f.maxX || y < f.minY || y > f.maxY {
		return false
	}

	if f.isVertical {
		return f.xvalue == x
	}
	return (f.m*x + f.c) == y
}

func newLinearFormula(x1 int, y1 int, x2 int, y2 int) LinearFormula {
	gradient := 0
	if (x2 - x1) == 0 {
		return LinearFormula{
			isVertical: true,
			xvalue: x1,
			minX: utils.Min(x1, x2),
			maxX: utils.Max(x1, x2),
			minY: utils.Min(y1, y2),
			maxY: utils.Max(y1, y2),
		}
	} else {
		gradient = (y2 - y1) / (x2 - x1)
	}
	yIntercept := y2 - (gradient * x2)
	return LinearFormula{
		m: gradient,
		c: yIntercept,
		isVertical: false,
		xvalue: 0,
		minX: utils.Min(x1, x2),
		maxX: utils.Max(x1, x2),
		minY: utils.Min(y1, y2),
		maxY: utils.Max(y1, y2),
	}
}

func AB(includeDiagonalLines bool, printOutput bool) {
	utils.SetBasePathToCurrentDir()

	// get the required file
	file, err := os.Open("input.txt")
	utils.HandleError(err, "opening input file")
	defer func(file *os.File) {
		err := file.Close()
		utils.HandleError(err, "closing input file")
	}(file)

	// stage it
	var lines []LinearFormula
	scanner := bufio.NewScanner(file)
	minX, minY, maxX, maxY := 0,0,0,0
	first := true
	for i := 0; scanner.Scan(); i++ {
		currentLine := strings.TrimSpace(scanner.Text())
		pairs := strings.Split(currentLine, " -> ")

		firstPair := strings.Split(pairs[0], ",")
		x1 := utils.String2Int(firstPair[0])
		y1 := utils.String2Int(firstPair[1])

		secondPair := strings.Split(pairs[1], ",")
		x2 := utils.String2Int(secondPair[0])
		y2 := utils.String2Int(secondPair[1])

		if includeDiagonalLines || (x1 == x2 || y1 == y2) {
			if first {
				maxX = utils.Max(x1, x2)
				maxY = utils.Max(y1, y2)

				minX = utils.Min(x1, x2)
				minY = utils.Min(y1, y2)

				first = false
			} else {
				if utils.Max(x1, x2) > maxX {
					maxX = utils.Max(x1, x2)
				}

				if utils.Max(y1, y2) > maxY {
					maxY = utils.Max(y1, y2)
				}

				if utils.Min(x1, x2) < minX {
					minX = utils.Min(x1, x2)
				}

				if utils.Min(y1, y2) < minY {
					minY = utils.Min(y1, y2)
				}
			}

			newFormula := newLinearFormula(x1, y1, x2, y2)
			if !contains(lines, newFormula) { // add it if it already doesnt exist
				lines = append(lines, newFormula)
			}
		}
	}

	// loop through all possible points and check how many points have >2 overlaps
	overlapCount := 0
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			overlapCountForPoint := 0
			for _, line := range lines {
				if line.isPointOnLine(x, y) {
					overlapCountForPoint++
				}
			}

			if printOutput {
				if overlapCountForPoint == 0 {
					fmt.Print(".")
				} else {
					fmt.Print(overlapCountForPoint)
				}
			}

			if overlapCountForPoint >= 2 {
				overlapCount++
			}
		}
		if printOutput {
			fmt.Println()
		}
	}

	fmt.Println("\nNumber of points where >=2 lines overlap: ", overlapCount)
}

func contains(ls []LinearFormula, el LinearFormula) bool {
	found := false
	for _, e := range ls {
		if e == el {
			found = true
			break
		}
	}
	return found
}
