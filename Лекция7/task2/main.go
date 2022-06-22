package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	input := "1 11 3 4 -5"
	var results string
	var min int32
	var max int32

	dividedInput := strings.Fields(input)

	for i, value := range dividedInput {
		v, _ := strconv.ParseInt(value, 10, 32)
		if i == 0 {
			max = int32(v)
			min = int32(v)
		}
		if int32(v) < min {
			min = int32(v)
		}
		if int32(v) > max {
			max = int32(v)
		}
	}
	minValue := strconv.FormatInt(int64(min), 10)
	maxValue := strconv.FormatInt(int64(max), 10)

	if max == min {
		results = maxValue
	} else {
		results = (maxValue + " " + minValue)
	}
	fmt.Println("results=", results)
}
