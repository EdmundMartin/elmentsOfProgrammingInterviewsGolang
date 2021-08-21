package stacks_queues

import (
	"strings"
)

var lookup = map[string]string{
	"(": ")",
	"{": "}",
	"[":"]",
}


func IsWellFormed(s string) bool  {
	leftChars := []string{}
	for _, char := range strings.Split(s, "") {
		if _, ok := lookup[char]; ok {
			leftChars = append(leftChars, char)
		} else if len(leftChars) == 0 {
			return false
		} else {
			x := leftChars[len(leftChars)-1]
			leftChars = leftChars[:len(leftChars) - 1]
			comparator, _ := lookup[x]
			if comparator != char {
				return false
			}
		}
	}
	return len(leftChars) == 0
}
