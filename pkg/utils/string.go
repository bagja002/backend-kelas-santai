package utils

import (
	"strconv"
	"strings"
)

func IntToString(x int) string {
	Number := strconv.Itoa(x)
	return Number
}
func UintToString(x uint64) string {
	Number := strconv.Itoa(int(x))
	return Number
}

func StringToInt(x string) int {
	Number, _ := strconv.Atoi(x)
	return Number
}

// StringToBool converts a string to a boolean.
// It returns true if the string is "true" (case-insensitive), and false otherwise.

func StringToBool(s string) bool {
	normalized := strings.TrimSpace(strings.ToLower(s))
	return normalized == "true" || normalized == "1" || normalized == "yes" || normalized == "y"
}

func StringToUint64(s string) uint64 {
	val, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0
	}
	return val
}

func RemoverSpaci(input string) string {

	remover := strings.ReplaceAll(input, " ", "")

	return remover
}
