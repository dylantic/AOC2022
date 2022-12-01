package Common

import (
	"sort"
	"strconv"
)

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

// SortIntSlice implements Go's standard sort library and makes reversing the slice a little quicker to implement.
// desc is a bool deciding whether the sort should be descending (true) or ascending (false)
func SortIntSlice(intSlice []int, desc bool) {
	if desc {
		sort.SliceStable(intSlice, func(i, j int) bool { return intSlice[i] > intSlice[j] })
	} else {
		sort.SliceStable(intSlice, func(i, j int) bool { return intSlice[i] < intSlice[j] })
	}
}
