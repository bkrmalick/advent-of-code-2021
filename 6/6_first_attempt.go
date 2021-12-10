package main

import (
	"advent-of-code-21/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func FirstAttempt(daysToCheck int) {
	utils.SetBasePathToCurrentDir()

	// get the required file
	file, err := os.Open("input.txt")
	utils.HandleError(err, "opening input file")
	defer func(file *os.File) {
		err := file.Close()
		utils.HandleError(err, "closing input file")
	}(file)

	// stage it
	x:= make([]*Fish, 0)
	population  := &x
	scanner := bufio.NewScanner(file)

	for i := 0; scanner.Scan(); i++ {
		currentLine := strings.TrimSpace(scanner.Text())

		fishies := strings.Split(currentLine, ",")

		for _,f := range fishies{
			newPopulation := append(*population, &Fish{utils.String2Int(f)})
			population = &newPopulation
		}
	}

	originalCount := len(*population)
	for day:=0;day<daysToCheck;day++{
		var newFishSpawned []*Fish
		for _,f:= range *population{
			newFish := f.passDay()
			if newFish != nil {
				newFishSpawned=append(newFishSpawned, newFish)
			}
		}
		tmp:=append(*population, newFishSpawned...)
		population = &tmp
	}

	fmt.Println("-----------------------------")
	fmt.Println("Original population size: ",originalCount)
	fmt.Println("New population size     : ",len(*population))
	fmt.Println("Difference              : ", len(*population)-originalCount)
}


