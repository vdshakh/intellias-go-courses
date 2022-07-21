package main

import (
	"fmt"
)

const (
	feedDog = 2
	feedCat = 7
	feedCow = 25
)

type isEatable map[string]struct{}

var isEatableList = isEatable{
	"Cow": {},
}

type Dog struct {
	weight int
}

func (d Dog) eatPerMonth() int {
	return d.weight * feedDog
}

func (d Dog) getWeight() int {
	return d.weight
}

func (d Dog) String() string {
	return "Dog"
}

type Cat struct {
	weight int
}

func (c Cat) eatPerMonth() int {
	return c.weight * feedCat
}

func (c Cat) getWeight() int {
	return c.weight
}

func (c Cat) String() string {
	return "Catt"
}

type Cow struct {
	weight int
}

func (cw Cow) eatPerMonth() int {
	return cw.weight * feedCow
}

func (cw Cow) getWeight() int {
	return cw.weight
}

func (cw Cow) String() string {
	return "Cow"
}

type Animal interface {
	eatPerMonth() int
	getWeight() int
	fmt.Stringer
}

// until better times
//type Farm struct {
//	cattle []Animal
//}
