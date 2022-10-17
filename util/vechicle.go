package util

import (
	"fmt"
)

// 底盘

type Chassis struct {
	Class int
}

func (c *Chassis) Rise() {
	fmt.Printf("Chassis %d rising.\n", c.Class)
}

func (c *Chassis) Lower() {
	fmt.Printf("Chassis %d lowering.\n", c.Class)
}

// 车身

type Body struct {
	Weight int
}

func (b *Body) Rise() {
	fmt.Printf("Body weighted-%d rising.\n", b.Weight)
}

func (b *Body) Lower() {
	fmt.Printf("Body weighted-%d lowering.\n", b.Weight)
}

// 引擎

type Engine struct {
	Brand string
}

func (e *Engine) Boot() {
	fmt.Printf("Engine %s boots.\n", e.Brand)
}

func (e *Engine) Halt() {
	fmt.Printf("Engine %s halts.\n", e.Brand)
}

// 方向盘

type Wheel struct {
	Diameter int
}

func (w *Wheel) TurnLeft() {
	fmt.Printf("Wheel with %d diameter turning left.\n", w.Diameter)
}

func (w *Wheel) TurnRight() {
	fmt.Printf("Wheel with %d diameter turning right.\n", w.Diameter)
}

// 车

type Vehicle struct {
	Chassis
	Body
	Engine
	Wheel
}

func (v *Vehicle) Run() {
	fmt.Println("Vehicle is running.")

	// v.Rise() // ambiguous selector v.Rise
	// v.Lower() // ambiguous selector v.Lower
	v.Chassis.Rise()  // OK
	v.Chassis.Lower() // OK
	v.Body.Rise()     // OK
	v.Body.Lower()    // OK

	v.Boot()
	v.Halt()
	v.TurnLeft()
	v.TurnRight()
}
