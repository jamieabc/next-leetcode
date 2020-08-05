package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func arguments() (*Difficulty, *int, error) {
	if len(os.Args) == 1 {
		helpMessage()
		return nil, nil, fmt.Errorf("insufficient arguments")
	}

	// parse arguments
	diffStr := flag.String("d", "", "difficulty: e(easy), m(medium), h(hard)")
	flag.Parse()

	problemNumber, err := strconv.Atoi(os.Args[len(os.Args)-1])
	if nil != err {
		fmt.Printf("input %s convert to integer with error: %s\n", os.Args[len(os.Args)-1], err)
		return nil, nil, fmt.Errorf("error problem number")
	}

	// parse difficulty
	diff := parseDifficulty(*diffStr)

	return &diff, &problemNumber, nil
}

func parseDifficulty(str string) Difficulty {
	if str == "" {
		return Difficulty{
			Easy:   true,
			Medium: true,
			Hard:   true,
		}
	}

	diff := Difficulty{}
	for i := range str {
		switch str[i] {
		case 'e':
			diff.Easy = true
		case 'm':
			diff.Medium = true
		case 'h':
			diff.Hard = true
		}
	}

	return diff
}

func helpMessage() {
	fmt.Println("find next problem: next-leecode $number")
}
