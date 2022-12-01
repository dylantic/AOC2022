package main

import (
	"fmt"
	Common "github.com/dylantic/AOC2022/common"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var data = string(input)
	var splitElves = strings.Split(data, "\r\n")
	elfMap := make(map[int]int)
	elf := 1

	for _, cal := range splitElves {
		//println(cal)
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

	sort.Sort(sort.Reverse(sort.IntSlice(calArray)))

	fmt.Printf("The elf carrying the most calories has %d in total\n", mostCalories)
	fmt.Printf("The top 3 elves are carrying %d in total\n", Common.Sum(calArray[0:3]))
}

func stringToInt(stringVar string) int {
	i, _ := strconv.Atoi(stringVar)
	return i
}

func sum(integers []int) int {
	result := 0
	for _, i := range integers {
		result += i
	}
	return result
}
