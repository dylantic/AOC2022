package main

import (
	"fmt"
	Common "github.com/dylantic/AOC2022/common"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var data = string(input)
	var splitElves = strings.Split(data, "\r\n")
	elfMap := make(map[int]int)
	elf := 1

	for _, cal := range splitElves {
		if len(cal) == 0 {
			elf = elf + 1
		} else {
			elfMap[elf] += Common.StringToInt(cal)
		}
	}

	var mostCalories = 0
	var calArray []int

	for _, calTotal := range elfMap {
		if calTotal > mostCalories {
			mostCalories = calTotal

		}
		calArray = append(calArray, calTotal)
	}

	// Weird-ass Go shenanigans to sort an array of integers in descending order.
	sort.Sort(sort.Reverse(sort.IntSlice(calArray)))

	fmt.Printf("The elf carrying the most calories has %d in total\n", mostCalories)
	fmt.Printf("The top 3 elves are carrying %d in total\n", Common.Sum(calArray[0:3]))
}
