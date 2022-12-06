package day2

import (
	"strings"
)

const (
	rock     = "rock"
	paper    = "paper"
	scissors = "scissors"
	win      = "win"
	lose     = "lose"
	draw     = "draw"
)

var OpponentMove map[string]string = map[string]string{
	"A": rock,
	"B": paper,
	"C": scissors,
}

var PlayerMove map[string]string = map[string]string{
	"X": rock,
	"Y": paper,
	"Z": scissors,
}

var PlayerDecision map[string]string = map[string]string{
	"X": lose,
	"Y": draw,
	"Z": win,
}

var MoveValues map[string]int = map[string]int{
	rock:     1,
	paper:    2,
	scissors: 3,
}

var OutcomeValue map[string]int = map[string]int{
	win:  6,
	lose: 0,
	draw: 3,
}

var OutcomePoints map[string]map[string]int = map[string]map[string]int{
	rock: {
		rock:     OutcomeValue[draw],
		paper:    OutcomeValue[win],
		scissors: OutcomeValue[lose],
	},
	paper: {
		rock:     OutcomeValue[lose],
		paper:    OutcomeValue[draw],
		scissors: OutcomeValue[win],
	},
	scissors: {
		rock:     OutcomeValue[win],
		paper:    OutcomeValue[lose],
		scissors: OutcomeValue[draw],
	},
}

var OutcomePointsV2 map[string]map[string]int = map[string]map[string]int{
	rock: {
		win:  MoveValues[paper],
		lose: MoveValues[scissors],
		draw: MoveValues[rock],
	},
	paper: {
		win:  MoveValues[scissors],
		lose: MoveValues[rock],
		draw: MoveValues[paper],
	},
	scissors: {
		win:  MoveValues[rock],
		lose: MoveValues[paper],
		draw: MoveValues[scissors],
	},
}

func Part1(inputs []string) int {
	pointTotal := 0

	for _, input := range inputs {
		moves := strings.Fields(input)
		opponent, player := OpponentMove[moves[0]], PlayerMove[moves[1]]
		pointTotal += OutcomePoints[opponent][player] + MoveValues[player]
	}

	return pointTotal
}

func Part2(inputs []string) int {
	pointTotal := 0

	for _, input := range inputs {
		moves := strings.Fields(input)
		opponent, player := OpponentMove[moves[0]], PlayerDecision[moves[1]]

		pointTotal += OutcomePointsV2[opponent][player] + OutcomeValue[player]
	}

	return pointTotal
}
