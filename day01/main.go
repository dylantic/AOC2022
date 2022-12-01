package main

import (
	"fmt"
	Common "github.com/dylantic/AOC2022/common"
	"log"
	"os"
	"strings"
)

func main() {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	data := string(input)
	splitElves := strings.Split(data, "\r\n")
	elfMap := make(map[int]int)
	elf := 1

	mostCalories := 0
	calSlice := make([]int, 0)

	for _, cal := range splitElves {
		if len(cal) == 0 {
			elf = elf + 1
		} else {
			elfMap[elf] += Common.StringToInt(cal)
		}
		if elfMap[elf] > mostCalories {
			mostCalories = elfMap[elf]
		}
		calSlice = append(calSlice, elfMap[elf])
	}

	Common.SortIntArray(calSlice, true)

	fmt.Printf("The elf carrying the most calories has %d in total\n", mostCalories)
	fmt.Printf("The top 3 elves are carrying %d in total\n", Common.Sum(calSlice[0:3]))
}
