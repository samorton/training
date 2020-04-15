package main

import (
	"fmt"
)

type Dog struct {
	Animal
}

type Animal struct {
	Age int
}

func (a *Animal) Move() {
	fmt.Printf("Animal moved\n")
}

func (a *Animal) SayAge() {
	fmt.Printf("Animal's age: %d", a.Age)
}

func main() {

	d := new(Dog)

	d.Age = 3

	d.Move()

	d.SayAge()

}
