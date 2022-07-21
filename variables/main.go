package main

import (
	"fmt"
	"math"
)

const money = 23.00

func main() {
	applePrice := 5.99
	pearPrice := 7.00
	numberOfApples := 9
	numberOfPears := 8
	// #1
	fmt.Printf("1. Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?")

	firstAnswer := (applePrice * float64(numberOfApples)) + (pearPrice * float64(numberOfPears))
	fmt.Printf("\nAmount of money spent: %v", firstAnswer)
	// #2
	fmt.Printf("\n2. Скільки груш ми можемо купити?")

	secondAnswer := math.Round(money / pearPrice)
	fmt.Printf("\nNumber of pears: %v", secondAnswer)
	// #3
	fmt.Printf("\n3. Скільки яблук ми можемо купити?")

	thirdAnswer := math.Round(money / applePrice)
	fmt.Printf("\nNumber of apples: %v", thirdAnswer)
	// #4
	fmt.Printf("\n4. Чи ми можемо купити 2 груші та 2 яблука?")

	quantity := 2
	fourthAnswer := money > (applePrice+pearPrice)*float64(quantity)
	fmt.Printf("\nThe statement \"I can buy 2 pears and 2 apples\" is %v", fourthAnswer)
}
