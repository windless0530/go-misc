package util

import (
	"fmt"
)

type Eatable interface {
	Eat()
}

type Person struct {
}

func (p *Person) Eat() {
	fmt.Println("I am person, amd I am eating.")
}

var _ Eatable = &Person{}

func RunEatable(e Eatable) {
	e.Eat()
}

func RunEatable2[E Eatable](e E) {
	e.Eat()
}
