package Common

import "strconv"

// StringToInt is a wrapper function for strconv.Atoi, returning the integer without errors.
func StringToInt(stringVar string) int {
	i, _ := strconv.Atoi(stringVar)
	return i
}

// Sum takes an array or a slice of integers and returns the sum.
func Sum(integers []int) int {
	result := 0
	for _, i := range integers {
		result += i
	}
	return result
}
