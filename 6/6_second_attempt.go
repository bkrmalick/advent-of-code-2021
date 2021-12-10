package main

import (
	"advent-of-code-21/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func SecondAttempt(daysToCheck int) {
	utils.SetBasePathToCurrentDir()

	// get the required file
	file, err := os.Open("input.txt")
	utils.HandleError(err, "opening input file")
	defer func(file *os.File) {
		err := file.Close()
		utils.HandleError(err, "closing input file")
	}(file)

	// stage it
	worldState := make(map[int]int) // <timer_value, count_of_fishes_with_this_value>
	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		currentLine := strings.TrimSpace(scanner.Text())
		fishTimers := strings.Split(currentLine, ",")
		for _,timerVal := range fishTimers{
			worldState[utils.String2Int(timerVal)]++
		}
	}

	const newFishTimer int = 8 // a new fish should be set to this timer value
	const matureFishTimer int = 6 // a mature fish should be reset to this timer value once it reaches 0

	// loop through the days updating the world state map as needed
	for day:=0;day<daysToCheck;day++ {
		originalState := cloneMap(worldState)
		// shift each count back a timer value
		for i:=0;i<newFishTimer;i++ {
			worldState[i] = originalState[i+1]
		}
		// only special case is what to do with counts at 0
		// 1. account for new fishes spawned
		worldState[newFishTimer]= originalState[0]
		// 2. reset existing fishes to 6
		worldState[matureFishTimer]+= originalState[0]
	}

	// loop through all possible timer values and sum them
	sum := 0
	for _,countOfFish := range worldState {
		sum+=countOfFish
	}

	// output
	fmt.Println("SUM: ", sum)
}

func cloneMap(originalMap map[int]int) map[int]int{
	newMap := make(map[int]int)
	for k,v := range originalMap {
		newMap[k] = v
	}
	return newMap
}


