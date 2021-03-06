package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Problem struct {
	Number     int
	Title      string
	Difficulty string
}

func problems() []Problem {
	result := make([]Problem, 0)

	lines := strings.Split(stats, "\n")
	for _, l := range lines {
		// comment starts with #
		if len(l) > 0 && l[0] != '#' && l != "\r" {
			// 123,.....
			words := strings.Split(l, ",")
			if len(words) <= 1 {
				// contains only \r
				continue
			}

			num, err := strconv.Atoi(words[0])
			if nil != err {
				fmt.Printf("parse line %v with error: %s\n", words, err)
				continue
			}

			result = append(result, Problem{
				Number:     num,
				Title:      words[1],
				Difficulty: words[3],
			})
		}
	}

	return result
}

// next - get next problem index by difficulty
func next(info []Problem, currentProblem *int, difficulty *Difficulty) (int, error) {
	var i int
	for i = range info {
		if info[i].Number == *currentProblem {
			i++
			break
		}
	}

	for i < len(info) {
		if difficulty.Match(info[i].Difficulty) {
			return i, nil
		}
		i++
	}

	return -1, fmt.Errorf("problem not found")
}
