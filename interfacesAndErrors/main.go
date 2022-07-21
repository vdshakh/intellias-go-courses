package main

import (
	"errors"
	"fmt"
)

func countFeed(a []Animal) int {
	var feedForAnimal int

	for _, animal := range a {
		fmt.Printf("%s weighing %v kg needs %v kg of feed\n", animal, animal.getWeight(), animal.eatPerMonth())
		feedForAnimal += animal.eatPerMonth()
	}

	return feedForAnimal
}

var (
	typeErr            = errors.New("typeAnimal doesn't match the animal name")
	unknownTypeErr     = errors.New("unknown animalType")
	incorrectWeightErr = errors.New("weight is incorrect")
	isEatableErr       = errors.New("animal isn't eatable")
)

func validator(a Animal) error {
	if err := validateType(a); err != nil {
		return fmt.Errorf("validateType failed: %w", err)
	}

	if err := validateWeight(a); err != nil {
		return fmt.Errorf("validateWeight failed: %w", err)
	}

	if err := validateIsEatable(a); err != nil {
		return fmt.Errorf("validateIsEatable failed: %w", err)
	}

	return nil
}

func validateType(a Animal) error {
	switch animalType := a.(type) {
	case Dog:
		if "Dog" != animalType.String() {
			return typeErr
		}
	case Cat:
		if "Cat" != animalType.String() {
			return typeErr
		}
	case Cow:
		if "Cow" != animalType.String() {
			return typeErr
		}
	default:
		return unknownTypeErr
	}

	return nil
}

func validateWeight(a Animal) error {
	minWeight := 5
	weight := a.getWeight()

	if weight < minWeight {
		return incorrectWeightErr
	}

	return nil
}

func validateIsEatable(a Animal) error {
	if _, ok := isEatableList[a.String()]; !ok {
		return isEatableErr
	}

	return nil
}

func main() {
	animals := []Animal{
		Dog{weight: 12},
		Cat{weight: 7},
		Cow{weight: 88},
		Dog{weight: 3},
	}

	feedPerMonth := countFeed(animals)
	fmt.Printf("\nKG OF FEED PER MONTH IS NEEDED: %v \n", feedPerMonth)

	for i, value := range animals {
		err := validator(animals[i])
		if err != nil {
			if errors.Is(err, incorrectWeightErr) {
				fmt.Printf("\n\nDon't validate for isEatable if validate for correct weight is failed ")
				continue
			}

			fmt.Printf("\n\nvalidate for %v failed: %w", value, err)
		}
	}
}
