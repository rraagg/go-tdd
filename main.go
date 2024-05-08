package main

import (
	"fmt"
)

func main() {
	fmt.Println("Doing go TDD")
	calc := InitCalc()
	calc.PrintAdd(2.5, 6.3)
}
