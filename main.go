package main

import (
	"fmt"

	"github.com/rraagg/go-tdd/calculator"
)

func main() {
	fmt.Println("Doing go TDD")
	engine := calculator.NewEngine()
	calc := calculator.NewCalculator(engine)
	calc.PrintAdd(2, 3)
}
