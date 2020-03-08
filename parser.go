package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	filename = "data.csv"
)

func parse() ([]int, error) {
	result := make([]int, 0)

	data, err := ioutil.ReadFile(filename)
	if nil != err {
		fmt.Printf("read file %s with error: %s\n", filename, err)
		return result, err
	}

	lines := strings.Split(string(data), "\n")
	for _, l := range lines {
		// comment starts with #
		if !strings.Contains(l, "#") && l != "\r" {
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
			result = append(result, num)
		}
	}

	return result, nil
}
