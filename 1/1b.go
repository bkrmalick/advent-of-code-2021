package main

import (
	"advent-of-code-21/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func B() {
	utils.SetBasePathToCurrentDir()

	file, err := os.Open("input.txt")

	utils.HandleError(err, "while opening input file")

	scanner := bufio.NewScanner(file)
	count := 0

	var prev int = 0
	var next int = 0
	var current int = 0

	scanner.Scan()
	current, _ = strconv.Atoi(scanner.Text())

	scanner.Scan()
	next, _ = strconv.Atoi(scanner.Text())

	more := scanner.Scan() // true if input greater than 3 lines
	for i := 0; more; i++ {
		prevWindowSum := prev + current + next

		prev = current
		current = next
		next, err = strconv.Atoi(scanner.Text())
		utils.HandleError(err, "trying to parse line to a int")

		currentWindowSum := prev + current + next
		if i != 0 && currentWindowSum > prevWindowSum {
			count++
		}
		more = scanner.Scan()
	}

	fmt.Println(count)
}
