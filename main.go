package main

import (
	"fmt"
	"os"
)

func main() {
	diff, currentProblem, err := arguments()
	if err != nil {
		os.Exit(1)
	}

	problemInfo := problems()

	idx, err := next(problemInfo, currentProblem, diff)

	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	fmt.Printf("Next problem of %d: %d (%s)\n", *currentProblem, problemInfo[idx].Number, problemInfo[idx].Difficulty)
}
