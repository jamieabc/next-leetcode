package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		help()
		os.Exit(0)
	}

	target, err := strconv.Atoi(os.Args[1])
	if nil != err {
		fmt.Printf("input %s convert to integer with error: %s\n", os.Args[1], err)
		os.Exit(0)
	}

	problems := parse()
	length := len(problems)

	for i, p := range problems {
		if p.Number == target {
			if i == length-1 {
				fmt.Println("Congratulations, this is the last problem")
				os.Exit(0)
			} else {
				fmt.Printf("Next problem of %d: %d (%s)\n", target, problems[i+1].Number, problems[i+1].Difficulty)
			}
		}
	}
}

func help() {
	fmt.Println("find next problem: next-leecode $number")
}
