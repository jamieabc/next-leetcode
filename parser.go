package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func parse(filePath string) error {
	result := make([]int, 0)

	data, err := ioutil.ReadFile(filePath)
	if nil != err {
		fmt.Printf("read file %s with error: %s\n", filePath, err)
		return err
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

	for _, n := range result {
		fmt.Printf("%d, ", n)
	}
	fmt.Println()

	return nil
}
