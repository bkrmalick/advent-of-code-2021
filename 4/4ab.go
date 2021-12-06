package main

import (
	"advent-of-code-21/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Point struct {
	val          int
	colNo, rowNo int
	marked       bool
}

type Board struct {
	points                 []*Point
	lastCalledNumber       int
	height, width          int
	won                    bool
}

func (b *Board) callNumber(pointVal int) {
	for _, p := range b.points {
		if p.val == pointVal {
			p.marked = true
			b.lastCalledNumber = pointVal

			if b.isRowFullyMarked(p.rowNo) || b.isColFullyMarked(p.colNo) {
				b.won = true
				return
			}
		}
	}

}

func (b *Board) isRowFullyMarked(rowNo int) bool {
	marked := true
	for _, p := range b.points {
		if p.rowNo == rowNo && !p.marked{
			marked = false
			break
		}
	}
	return marked
}

func (b *Board) isColFullyMarked(colNo int) bool {
	marked := true
	for _, p := range b.points {
		if p.colNo == colNo && !p.marked{
			marked = false
			break
		}
	}
	return marked
}


func (b *Board) calculateScore() int {
	sumOfUnmarked := 0

	for _, p := range b.points {
		if !p.marked {
			sumOfUnmarked += p.val
		}
	}

	return sumOfUnmarked * b.lastCalledNumber

}

func newBoard(lines []string) *Board {
	board := Board{points: make([]*Point, 0)}
	board.height = len(lines)

	colIndex := 0
	boardWidth := 0

	for rowIndex, row := range lines {
		boardWidth = 0
		cols := strings.Split(row, " ")
		for _, col := range cols {
			if col != "" {
				board.points = append(board.points,
					&Point{
						utils.String2Int(col),
						boardWidth,
						rowIndex,
						false,
					},
				)
				boardWidth++
			}
			colIndex++
		}
		colIndex = 0
	}

	board.width = boardWidth

	return &board
}

func AB() {
	utils.SetBasePathToCurrentDir()

	// get the required files
	numbersDrawnFile, err := os.Open("input_numbers_drawn.txt")
	utils.HandleError(err, "opening input file")
	defer func(file *os.File) {
		err := file.Close()
		utils.HandleError(err, "closing input file")
	}(numbersDrawnFile)

	boardsFile, err := os.Open("input_boards.txt")
	utils.HandleError(err, "opening input file")
	defer func(file *os.File) {
		err := file.Close()
		utils.HandleError(err, "closing input file")
	}(boardsFile)

	// dispatch the bingo subsystem
	numbersDrawnChan := make(chan int)
	go runBingoSubsystem(numbersDrawnFile, numbersDrawnChan)

	var boards []*Board
	var boardLines []string

	scanner := bufio.NewScanner(boardsFile)
	for scanner.Scan() {
		currentLine := strings.TrimSpace(scanner.Text())

		if currentLine != "" {
			boardLines = append(boardLines, currentLine)
		} else {
			// one board has ended, next one about to begin
			boards = append(boards, newBoard(boardLines))
			boardLines = nil
		}
	}

	boards = append(boards, newBoard(boardLines))

	for calledNum := range numbersDrawnChan {
		for boardIndex, b := range boards {
			alreadyWon := b.won
			b.callNumber(calledNum)
			if b.won && !alreadyWon {
				fmt.Println("Board number ", boardIndex, " won with score ", b.calculateScore())
			}

		}
	}
}

// runs a goroutine that will dispatch the drawn numbers
// to the provided channel
func runBingoSubsystem(file *os.File, numbersDrawn chan int) {
	lines := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentLine := strings.TrimSpace(scanner.Text())
		lines = append(lines, currentLine)
	}

	if len(lines) > 1 {
		log.Fatalf("Expected only one line for the numbers drawn input file")
	}

	inputNums := strings.Split(lines[0], ",")

	for _, v := range inputNums {
		numbersDrawn <- utils.String2Int(v)
	}

	close(numbersDrawn)
}
