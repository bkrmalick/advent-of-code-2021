package main

import (
	"advent-of-code-21/utils"
	"bufio"
	"fmt"
	"github.com/montanaflynn/stats"
	"log"
	"math"
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
	hPositions := make([]int, 0)
	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		currentLine := strings.TrimSpace(scanner.Text())
		poss := strings.Split(currentLine, ",")
		for _, p := range poss {
			hPositions = append(hPositions, utils.String2Int(p))
		}
	}

	optimalPosition1 := calculateMedian(hPositions)
	fuelCostConstant := calculateConstFuelCost(hPositions, optimalPosition1)

	optimalPosition21,optimalPosition22 := calculateMean(hPositions)
	fuelCostVariable21, fuelCostVariable22 := calculateVarFuelCost(hPositions, optimalPosition21), calculateVarFuelCost(hPositions, optimalPosition22)

	fmt.Println("Solution A: Optimal position is", optimalPosition1, "which will cost", fuelCostConstant)
	fmt.Println("Solution B: Optimal positions are (", optimalPosition21,",",optimalPosition22,") which will cost (", fuelCostVariable21, ",",fuelCostVariable22,")")
}

func calculateConstFuelCost(hPositions []int, optimalPosition int) int {
	fuelCost := 0
	for _, p := range hPositions {
		diff := p - optimalPosition
		if diff < 0 {
			diff *= -1
		}
		fuelCost += diff
	}
	return fuelCost
}

func calculateVarFuelCost(hPositions []int, optimalPosition int) int {
	fuelCost := 0
	for _, p := range hPositions {
		diff := p - optimalPosition
		if diff < 0 {
			diff *= -1
		}
		fuelCost += (diff * (diff + 1)) / 2 // formula to get 1+2+...+n sum
	}
	return fuelCost
}

func calculateMedian(positions []int) int {
	median, err := stats.Median(stats.LoadRawData(positions))
	if err != nil {
		log.Fatalf("Error while finding median: %s", err)
	}
	return int(median + 0.5)
}

//	mean is highly affected by outliers which we DO want to prioritise since it will cost
//	outlying crabs more fuel per unit distance

// 	returns roundedDown. roundedUp values
func calculateMean(positions []int) (int, int) {
	mean, err := stats.Mean(stats.LoadRawData(positions))
	if err != nil {
		log.Fatalf("Error while finding finding mean: %s", err)
	}
	return int(math.Floor(mean)), int(math.Ceil(mean))  // round to the nearest int
}
