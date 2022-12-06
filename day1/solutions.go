package day1

import (
	"sort"
	"strconv"

	"github.com/FrankMoreno/adventcode_2022/utils"
)

const (
	EMPTY_LINE = ""
)

func Part1(inputs []string) int {
	localMax, globalMax := 0, 0

	for _, input := range inputs {
		if input == "" {
			globalMax = utils.Max(localMax, globalMax)
			localMax = 0
		} else {
			food, _ := strconv.Atoi(input)
			localMax += food
		}
	}

	return globalMax
}

func Part2(inputs []string) int {
	localMax := 0
	topThree := []int{}

	for _, input := range inputs {
		if input == EMPTY_LINE {
			if len(topThree) < 3 {
				topThree = append(topThree, localMax)
				sort.Ints(topThree)
			} else if localMax > topThree[0] {
				topThree[0] = localMax
				sort.Ints(topThree)
			}
			localMax = 0
		} else {
			food, _ := strconv.Atoi(input)
			localMax += food
		}
	}

	return utils.Sum(topThree)
}

func Part2v2(inputs []string) int {
	localMax := 0
	allTotals := []int{}

	for _, input := range inputs {
		if input == EMPTY_LINE {
			allTotals = append(allTotals, localMax)
			localMax = 0
		} else {
			food, _ := strconv.Atoi(input)
			localMax += food
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(allTotals)))

	return utils.Sum(allTotals[0:3])
}
