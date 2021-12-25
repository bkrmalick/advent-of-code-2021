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

	// The only difference between A and B parts of the problem is the predicate which decides
	// which adjacent node can be visited
	predicateA := func(nextNodeToVisit string, currentPath []string) bool {
		return strings.ToUpper(nextNodeToVisit) == nextNodeToVisit || !utils.InList(nextNodeToVisit, currentPath)
	}
	predicateB := func(nextNodeToVisit string, currentPath []string) bool {
		return strings.ToUpper(nextNodeToVisit) == nextNodeToVisit || !utils.InList(nextNodeToVisit, currentPath) || canVisitSmallCaveTwice(currentPath, nextNodeToVisit)
	}

	solutionsCountA := getSolutionCountWithPredicate(adjList, predicateA)
	solutionsCountB := getSolutionCountWithPredicate(adjList, predicateB)

	fmt.Println("\n (A) Total count of solutions : ", solutionsCountA)
	fmt.Println("\n (B) Total count of solutions : ", solutionsCountB)
}

func getSolutionCountWithPredicate(
	adjList map[string][]string,
	predicate func(nextNodeToVisit string, currentPath []string) bool, ) int {

	// channel which will recv solutions from goroutines
	// each solution will be a string of the form "start, A, B, C, end"
	solutions := make(chan string)

	// explore all possible paths concurrently
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go explore(adjList, "start", []string{}, solutions, wg, predicate)

	// run a separate goroutine that will wait for all goroutines to finish and only then close the channel
	go utils.CloseChannelWhenDone(wg, solutions)

	return countSolutions(solutions)
}

func countSolutions(solutions chan string) int {
	count := 0
	for sol := range solutions {
		fmt.Println(sol)
		count++
	}
	return count
}

func explore(
	adjList map[string][]string,
	currentNode string,
	currentPath []string,
	finalSolutions chan string, // channel which records all solutions from all goroutines
	wg *sync.WaitGroup,
	predicate func(nextNodeToVisit string, currentPath []string) bool, // predicate which decides which adjacent nodes to visit
) {
	defer wg.Done()

	if currentNode == "end" {
		// if we're at the end node, we've found a solution - send it to the channel as a csv string
		finalSolutions <- utils.StringSliceToCSV(append(currentPath, currentNode))
		return
	}

	nextPossibleNodes := adjList[currentNode]

	for _, n := range nextPossibleNodes {
		// clone array to provide each goroutine with its independent copy
		newPath := make([]string, len(currentPath))
		copy(newPath, currentPath)
		// add current node to the path
		newPath = append(newPath, currentNode)

		if predicate(n, newPath) {
			// dispatch a goroutine for each possible next node
			wg.Add(1)
			go explore(adjList, n, newPath, finalSolutions, wg, predicate)
		}
	}
}

func canVisitSmallCaveTwice(alreadyVisitedCaves []string, wantToVisitNode string) bool {
	freqMap := make(map[string]int)

	for _, c := range alreadyVisitedCaves {
		freqMap[c] ++
	}

	for _, c := range alreadyVisitedCaves {
		if strings.ToLower(c) == c && freqMap[c] == 2 { // cant visit node if there's already a small cave node that's been visited twice
			return false
		}
		if utils.InList(wantToVisitNode, []string{"start", "end"}) && freqMap[c] == 1 { // can't visit start and end more than once
			return false
		}
	}

	return true
}

func addRelationshipToAdjList(adjList map[string][]string, srcNode string, tgtNode string) {
	if adjList[srcNode] == nil {
		adjList[srcNode] = make([]string, 0)
	}
	if !utils.InList(tgtNode, adjList[srcNode]) {
		adjList[srcNode] = append(adjList[srcNode], tgtNode)
	}
}
