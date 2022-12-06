package main

import (
	"fmt"
	"log"
	"os"
	"unicode"

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

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(inputs []string) int {
	total := 0

	for _, input := range inputs {
		first_bag := input[:len(input)/2]
		second_bag := input[len(input)/2:]
		unique := map[rune]interface{}{}
		for _, letter := range first_bag {
			if _, ok := unique[letter]; !ok {
				unique[letter] = nil
			}
		}

		for _, letter := range second_bag {
			if _, ok := unique[letter]; ok {
				total += convertToPoints(letter)
				break
			}
		}
	}
	return total
}

func part2(inputs []string) int {
	total := 0
	for i, j := 0, 3; j <= len(inputs); i, j = i+3, j+3 {
		total += common3(inputs[i:j])
	}

	return total
}

func common3(inputs []string) int {
	unique := collections.NewSet[rune]()
	unique2 := collections.NewSet[rune]()

	for _, letter := range inputs[0] {
		if !unique.Contains(letter) {
			unique.Add(letter)
		}
	}

	for _, letter := range inputs[1] {
		if unique.Contains(letter) {
			unique2.Add(letter)
		}
	}

	for _, letter := range inputs[2] {
		if unique2.Contains(letter) {
			return convertToPoints(letter)
		}
	}

	return -1
}

func convertToPoints(c rune) int {
	if unicode.IsUpper(c) {
		return int(c) - 38
	} else {
		return int(c) - 96
	}
}
