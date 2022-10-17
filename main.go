package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// 1. 泛型声明和应用

func Run(e Eatable) {
	e.Eat()
}

func Calculate[T constraints.Ordered](a T, b T) T {
	return a + b
}

// 2. 自定义接口和实现

type Eatable interface {
	Eat()
}

type Person struct {
}

func (p *Person) Eat() {
	fmt.Println("I am person, amd I am eating.")
}

var _ Eatable = &Person{}

// main

func main() {
	fmt.Printf("%d\n", Calculate(1, 2))
	fmt.Printf("%.2f\n", Calculate(1.1, 2.3))

	a := Person{}
	Run(&a)
}
