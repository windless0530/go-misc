package main

import (
	"fmt"

	"golang.org/x/exp/constraints"

	"go-misc/util"
)

// 1. 泛型声明和应用

func Calculate[T constraints.Ordered](a T, b T) T {
	return a + b
}

func testGenericType() {
	fmt.Printf("%d\n", Calculate(1, 2))
	fmt.Printf("%.2f\n", Calculate(1.1, 2.3))
}

func testEatable() {
	a := util.Person{}
	util.RunEatable(&a)
	util.RunEatable2(&a)
}

func testVechicle() {
	a := util.Vehicle{
		Chassis: util.Chassis{Class: 1},
		Body:    util.Body{Weight: 1000},
		Engine:  util.Engine{Brand: "BMW"},
		Wheel:   util.Wheel{Diameter: 20},
	}
	a.Run()
}

func main() {
	testVechicle()
}
