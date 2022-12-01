package Common

import "strconv"

func StringToInt(stringVar string) int {
	i, _ := strconv.Atoi(stringVar)
	return i
}

func Sum(integers []int) int {
	result := 0
	for _, i := range integers {
		result += i
	}
	return result
}
