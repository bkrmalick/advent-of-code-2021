package main

import (
	"advent-of-code-21/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func A() {
	utils.SetBasePathToCurrentDir()

	file, err := os.Open("input.txt")

	utils.HandleError(err, "while opening input file")

	scanner := bufio.NewScanner(file)
	count := 0
	var prev int = 0
	for i:=0; scanner.Scan();i++ {
		var current int
		current, err := strconv.Atoi(scanner.Text())
		utils.HandleError(err, "trying to parse line to a int")
		if i != 0 && current>prev{
			count++
		}
		prev = current
	}

	fmt.Println(count)
}



