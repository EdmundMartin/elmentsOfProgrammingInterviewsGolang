package stacks_queues

import (
	"strconv"
	"strings"
)

var operators = map[string]func(y, x int) int{
	"+": func(y, x int) int {
		return x + y
	},
	"*": func(y, x int) int {
		return x * y
	},
	"-": func(y, x int) int {
		return x - y
	},
	"/": func(y, x int) int {
		return x / y
	},
}

func EvaluateRPN(expression string) int {
	intermediateResults := []int{}
	delimiter := ","

	for _, token := range strings.Split(expression, delimiter) {
		f, ok := operators[token]
		if ok {
			firstArg, x := intermediateResults[len(intermediateResults)-1], intermediateResults[:len(intermediateResults)-1]
			secondArg, y := x[len(x)-1], x[:len(x)-1]
			intermediateResults = y
			result := f(firstArg, secondArg)
			intermediateResults = append(intermediateResults, result)
		} else {
			val, _ := strconv.Atoi(token)
			intermediateResults = append(intermediateResults, val)
		}
	}
	return intermediateResults[len(intermediateResults)-1]
}
