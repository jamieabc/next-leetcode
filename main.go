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

	problems, err := parse()
	if nil != err {
		fmt.Println("parse problem set with error: ", err)
		os.Exit(0)
	}

	length := len(problems)

	for i, p := range problems {
		if p == target {
			if i == length-1 {
				fmt.Println("Congratulations, this is the last problem")
				os.Exit(0)
			} else {
				fmt.Printf("Next problem of %d is %d\n", target, problems[i+1])
			}
		}
	}
}

func help() {
	fmt.Println("Usage: next_leecode number")
}
