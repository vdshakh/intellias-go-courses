package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := "1 9 3 4 -5"
	var result string
	var max int
	var min int

	inputSplit := strings.Split(input, " ")
	inputArray := make([]int, len(inputSplit))

	max, _ = strconv.Atoi(inputSplit[0])
	min, _ = strconv.Atoi(inputSplit[0])

	for key, value := range inputSplit {
		inputArray[key], _ = strconv.Atoi(value)

		if max < inputArray[key] {
			max = inputArray[key]
		}

		if min > inputArray[key] {
			min = inputArray[key]
		}
	}

	result = strconv.Itoa(max) + " " + strconv.Itoa(min)

	fmt.Printf("%v\n", result)
}
