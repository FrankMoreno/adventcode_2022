package main

import (
	"fmt"
	"log"
	"os"

	"github.com/FrankMoreno/adventcode_2022/day6"
	"github.com/FrankMoreno/adventcode_2022/utils"
)

func main() {
	if len(os.Args) < 2 {
		log.Panic("NO FILE")
	}

	input, err := utils.ReadInput(os.Args[1], "\n")
	if err != nil {
		log.Panicf("Unable to read file contents: %v", err)
	}

	fmt.Println(day6.Part1(input[0]))
}
