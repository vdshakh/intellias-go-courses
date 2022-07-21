package main

import (
	"fmt"
	"sync"
)

// Конкурентно порахувати суму кожного слайсу int, та роздрукувати результат.
// Потрібно використовувати WaitGroup.
// Приклад:
// [ [ 4, 6 ], [ 7, 9 ] ]
// Результат друку:
// Порядок друку не важливий.
// “slice 1: 10”
// “slice 2: 16”
func main() {
	var wg sync.WaitGroup
	n := [][]int{
		{2, 6, 9, 24},
		{7, 3, 94, 3, 0},
		{4, 2, 8, 35},
	}

	for key, _ := range n {
		wg.Add(1)

		go func(m int) {
			defer wg.Done()
			fmt.Printf("slice %v: ", key+1) //added +1 for prettier printing
			sum(n[key])
		}(key)

		wg.Wait()
	}
}

func sum(input []int) {
	var result int

	for _, value := range input {
		result += value
	}

	fmt.Printf("%v\n", result)
}
