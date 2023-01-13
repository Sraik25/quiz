package utils

import (
	"fmt"
	"os"
	"strings"
)

func ParseLines(lines [][]string) []Problem {
	ret := make([]Problem, len(lines))

	for i, line := range lines {
		ret[i] = Problem{
			Q: line[0],
			A: strings.TrimSpace(line[1]),
		}
	}

	return ret
}

type Problem struct {
	Q string
	A string
}

func Exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
