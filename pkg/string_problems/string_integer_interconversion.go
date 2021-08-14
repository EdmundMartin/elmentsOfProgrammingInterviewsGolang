package string_problems

import (
	"bytes"
	"strings"
)

var simpleMap = map[string]int{
	"0": 0,
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
}

func IntToString(x int32) string {
	isNegative := false
	if x < 0 {
		x, isNegative = -x, true
	}
	s := []rune{}

	for {
		s = append(s, rune('0')+x%10)
		x /= 10
		if x == 0 {
			break
		}
	}

	var out bytes.Buffer
	if isNegative {
		out.WriteString("-")
	}
	for i := len(s) - 1; i >= 0; i-- {
		out.WriteString(string(s[i]))
	}

	return out.String()
}

func StrToInt(x string) int {
	start := 1
	if x[0] == '-' {
		start = -1
	}
	runningSum := 0
	for _, chr := range strings.Split(x, "") {
		if chr == "-" {
			continue
		} else {
			val, _ := simpleMap[chr]
			runningSum = runningSum*10 + val
		}
	}
	return start * runningSum
}
