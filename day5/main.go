package main

import (
	"fmt"
	"log"
	"os"

	"github.com/FrankMoreno/adventcode_2022/collections"
	"github.com/FrankMoreno/adventcode_2022/utils"
)

func main() {
	if len(os.Args) < 2 {
		log.Panic("NO FILE")
	}

	input, err := utils.ReadInput(os.Args[1], "\n")
	if err != nil {
		log.Panicf("Unable to read file contens: %v", err)
	}

	fmt.Println(part1(input, createRealInput()))
	fmt.Println(part2(input, createRealInput()))
}

func part1(inputs []string, crates []collections.Stack[string]) string {
	for index := 0; index < len(inputs); index++ {
		count, source, destination := parseInstruction(inputs[index])
		for i := 0; i < count; i++ {
			crates[destination].Push(crates[source].Pop())
		}
	}

	var answer string
	for _, crate := range crates {
		answer += crate.Pop()
	}

	return answer
}

func part2(inputs []string, crates []collections.Stack[string]) string {
	for index := 0; index < len(inputs); index++ {
		count, source, destination := parseInstruction(inputs[index])

		tempstack := collections.NewStack[string]()
		for i := 0; i < count; i++ {
			temp := crates[source].Pop()
			tempstack.Push(temp)
		}

		for i := 0; i < count; i++ {
			crates[destination].Push(tempstack.Pop())
		}
	}

	var answer string
	for _, crate := range crates {
		answer += crate.Pop()
	}

	return answer
}

func parseInstruction(input string) (int, int, int) {
	var count, source, destination int
	fmt.Sscanf(input, "move %d from %d to %d", &count, &source, &destination)
	return count, source - 1, destination - 1
}

func createTestInput() []collections.Stack[string] {
	test_input := make([]collections.Stack[string], 3)
	test_input[0] = collections.NewStack("Z", "N")
	test_input[1] = collections.NewStack("M", "C", "D")
	test_input[2] = collections.NewStack("P")

	return test_input
}

func createRealInput() []collections.Stack[string] {
	input := make([]collections.Stack[string], 9)
	input[0] = collections.NewStack("S", "T", "H", "F", "W", "R")
	input[1] = collections.NewStack("S", "G", "D", "Q", "H")
	input[2] = collections.NewStack("B", "T", "Q")
	input[3] = collections.NewStack("D", "R", "W", "T", "N", "Q", "Z", "J")
	input[4] = collections.NewStack("F", "B", "H", "G", "L", "V", "T", "Z")
	input[5] = collections.NewStack("L", "P", "T", "C", "V", "B", "S", "G")
	input[6] = collections.NewStack("Z", "B", "R", "T", "W", "G", "P")
	input[7] = collections.NewStack("N", "G", "M", "T", "C", "J", "R")
	input[8] = collections.NewStack("L", "G", "B", "W")

	return input
}
