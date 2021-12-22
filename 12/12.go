package main

import (
	"advent-of-code-21/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
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

	// stage it into a adjacency list
	lines := make([]string, 0)
	adjList := make(map[string][]string)
	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		currentLine := strings.TrimSpace(scanner.Text())
		splits := strings.Split(currentLine, "-")
		srcNode := splits[0]
		tgtNode := splits[1]

		addRelationshipToAdjList(adjList, srcNode, tgtNode)
		addRelationshipToAdjList(adjList, tgtNode, srcNode)

		lines = append(lines, currentLine)
	}

	// channels which will recv solutions from goroutines
	solutions := make(chan string)

	// explore all possible paths
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go explore(adjList, "start", []string{}, solutions, wg)

	// run a separate goroutine that will wait for all goroutines to finish and only then close the channel
	go utils.CloseChannelWhenDone(wg, solutions)

	countSolutions := 0
	for sol := range solutions {
		fmt.Println(sol)
		countSolutions++
	}

	fmt.Println("\nTotal count of solutions: ", countSolutions)
}

func explore(adjList map[string][]string, currentNode string, currentPath []string, finalSolutions chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	if currentNode == "end" {
		// if we're at the end node, we've found a solution - send it to the channel as a csv string
		finalSolutions <-utils.StringSliceToCSV(append(currentPath, currentNode))
		return
	}

	nextPossibleNodes := adjList[currentNode]

	for _, n := range nextPossibleNodes {
		if strings.ToUpper(n) == n || !utils.InList(n, currentPath) {
			// clone array to provide each goroutine with its independent copy
			newPath := make([]string, len(currentPath))
			copy(newPath, currentPath)
			newPath = append(newPath, currentNode)

			// dispatch a goroutine for each possible next node
			wg.Add(1)
			go explore(adjList, n,newPath, finalSolutions, wg)
		}
	}
}

func addRelationshipToAdjList(adjList map[string][]string, srcNode string, tgtNode string) {
	if adjList[srcNode] == nil {
		adjList[srcNode] = make([]string, 0)
	}
	if !utils.InList(tgtNode,adjList[srcNode]) {
		adjList[srcNode] = append(adjList[srcNode], tgtNode)
	}
}
