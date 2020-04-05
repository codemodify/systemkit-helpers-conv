package helpers

import (
	"strconv"
	"strings"
)

// IsNullOrEmpty -
func IsNullOrEmpty(value string) bool {
	return (len(strings.TrimSpace(value)) <= 0)
}

// Contains -
func Contains(a string, x string) bool {
	return (strings.Index(a, x) != -1)
}

// ArrayContains -
func ArrayContains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}

	return false
}

// StringToFloat64 -
func StringToFloat64(value string, returnIfError float64) float64 {
	result, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return returnIfError
	}

	return result
}

// StringToInt -
func StringToInt(value string, returnIfError int) int {
	result, err := strconv.Atoi(value)
	if err != nil {
		return returnIfError
	}

	return result

}

// StringToInt64 -
func StringToInt64(value string, returnIfError int64) int64 {
	result, err := strconv.Atoi(value)
	if err != nil {
		return returnIfError
	}

	return int64(result)

}
