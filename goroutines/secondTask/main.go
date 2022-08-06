package main

import "fmt"

// Конкурентно порахувати суму усіх слайсів int, та роздрукувати результат.
// Приклад:
// [ [ 4, 6 ], [ 7, 9 ] ]
// Результат друку:
// “result: 26”
func main() {
	n := [][]int{
		{2, 6, 9, 24},
		{7, 3, 94, 3, 0},
		{4, 2, 8, 35},
	}

	// Ваша реалізація
	ch := make(chan int)

	for i := 0; i < len(n); i++ {
		go sum(n[i], ch)
	}

	x, y, z := <-ch, <-ch, <-ch
	fmt.Println("Result: ", x+y+z)
}

func sum(s []int, c chan int) {
	result := 0

	for _, v := range s {
		result += v
	}

	c <- result
}
