package main

import (
	"strconv"
	"strings"
)

func isIPV4(addr string) bool {
	return false
}

func solution(addr string) bool {
	substrs := strings.Split(addr, ".")
	if len(substrs) != 4 {
		return false
	}
	for _, substr := range substrs {
		if len(substr) == 0 {
			return false
		}
		if len(substr) > 1 && substr[0] == '0' {
			return false
		}
		val, err := strconv.Atoi(substr)
		if err != nil {
			return false
		}
		if val < 0 || val > 255 {
			return false
		}
	}
	return true
}

func main() {}
