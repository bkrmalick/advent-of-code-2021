package main

import (
	"advent-of-code-21/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
)

type PositionWithAim struct {
	Position
	aim int
}

func (p *PositionWithAim) Up(n int){
	p.aim-=n
}

func (p *PositionWithAim) Down(n int){
	p.aim+=n
}

func (p *PositionWithAim) Forward(n int){
	p.Position.Forward(n) // do whatever we were doing previously in embedded struct
	p.Position.Down(p.aim*n) // then increase depth by aim * n
}

func (p *PositionWithAim) GetHorizontal() int {
	return p.horizontal
}

func (p *PositionWithAim) GetDepth() int {
	return p.depth
}


func B() {
	utils.SetBasePathToCurrentDir()

	file, err := os.Open("input.txt")

	utils.HandleError(err, "while opening input file")

	pos := PositionWithAim{}

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

	fmt.Println("Final position 2: ", pos)
	fmt.Println("Ans 2: ", pos.GetHorizontal() * pos.GetDepth())
}

