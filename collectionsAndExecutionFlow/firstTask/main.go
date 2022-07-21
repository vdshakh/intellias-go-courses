package main

import (
	"fmt"
)

func main() {
	arr := []int{4, 1, 4, -4, 6, 3, 8, 8}
	var result []int

	uniqElements := make(map[int]bool)

	for _, number := range arr {
		if _, value := uniqElements[number]; !value {
			uniqElements[number] = true
			result = append(result, number)
		}
	}

	fmt.Printf("%v\n", result)
}
