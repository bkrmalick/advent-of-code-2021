package main

import (
	"advent-of-code-21/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type Positioner interface {
	Up(n int)
	Down(n int)
	Forward(n int)
	GetHorizontal() int
	GetDepth() int
}

type Position struct {
	horizontal, depth int
}

func (p *Position) Up(n int){
	p.depth-=n
}

func (p *Position) Down(n int){
	p.depth+=n
}

func (p *Position) Forward(n int){
	p.horizontal+=n
}

func (p *Position) GetHorizontal() int {
	return p.horizontal
}

func (p *Position) GetDepth() int {
	return p.depth
}


func A() {
	utils.SetBasePathToCurrentDir()

	file, err := os.Open("input.txt")

	utils.HandleError(err, "while opening input file")

	pos := Position{}

	scanner := bufio.NewScanner(file)
	for i:=0; scanner.Scan();i++ {
		currentLine  := scanner.Text()
		operation, operand := parseOperationAndOperand(currentLine) // up, down, forward
		operation = strings.Title(strings.ToLower(operation))

		meth := reflect.ValueOf(&pos).MethodByName(operation) // use reflection to look for a method for the same name

		if !meth.IsValid() {
			log.Fatalf("Method %s not implemented yet on Position struct", operation)
		}

		meth.Call([]reflect.Value{reflect.ValueOf(operand)})

	}

	fmt.Println("Final position 1: ", pos)
	fmt.Println("Ans 1: ", pos.GetHorizontal() * pos.GetDepth())
}

func parseOperationAndOperand(s string) (string, int){
	newS := ""

	// loop through characters in string
	for index, c := range s {
		// operation name is up until an int is discovered
		if isInt,operand:=isStringInt(s[index:]); isInt {
			return strings.TrimSpace(newS), operand
		}
		newS+=string(c)
	}

	return "",0
}

func isStringInt(s string) (bool, int){
	if x, err := strconv.Atoi(strings.TrimSpace(s)); err == nil {
		return true,x
	}
	return false, 0
}
