package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/FrankMoreno/adventcode_2022/utils"
)

type Bounds struct {
	lower int
	upper int
}

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
		splitInput := strings.Split(input, ",")
		firstBounds, secondBounds := convertToBounds(splitInput[0]), convertToBounds(splitInput[1])

		if encompassing(firstBounds, secondBounds) || encompassing(secondBounds, firstBounds) {
			total += 1
		}
	}

	return total
}

func part2(inputs []string) int {
	total := 0
	for _, input := range inputs {
		splitInput := strings.Split(input, ",")
		firstBounds, secondBounds := convertToBounds(splitInput[0]), convertToBounds(splitInput[1])

		if overlapping(firstBounds, secondBounds) || overlapping(secondBounds, firstBounds) {
			total += 1
		}
	}

	return total
}

func encompassing(first, second Bounds) bool {
	if first.lower <= second.lower && first.upper >= second.upper {
		return true
	} else {
		return false
	}
}

func overlapping(first, second Bounds) bool {
	if second.lower >= first.lower && second.lower <= first.upper {
		return true
	} else if second.upper >= first.lower && second.upper <= first.upper {
		return true
	} else {
		return false
	}
}

func convertToBounds(input string) Bounds {
	ranges := strings.Split(input, "-")
	lower, _ := strconv.Atoi(ranges[0])
	upper, _ := strconv.Atoi(ranges[1])

	return Bounds{
		lower: lower,
		upper: upper,
	}
}
