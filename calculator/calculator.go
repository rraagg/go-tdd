package calculator

import "fmt"

type Adder interface {
	Add(x, y float64) float64
}

type Calculator struct {
	Adder Adder
}

func NewCalculator(a Adder) *Calculator {
	return &Calculator{Adder: a}
}

func (c Calculator) PrintAdd(x, y float64) {
	fmt.Println("Result:", c.Adder.Add(x, y))
}
