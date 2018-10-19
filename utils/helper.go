package utils

import (
	"strconv"
)

// convert string into integer
func StringToInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}

	return num
}
