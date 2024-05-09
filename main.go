package main

import (
	"flag"
	"log"

	"github.com/rraagg/go-tdd/calculator"
	"github.com/rraagg/go-tdd/input"
)

func main() {
	expr := flag.String("expression", "", "mathematical expression to parse")
	flag.Parse()

	engine := calculator.NewEngine()
	validator := input.NewValidator(engine.GetNumOperands(), engine.GetValidOperators())
	parser := input.NewParser(engine, validator)
	result, err := parser.ProcessExpression(*expr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(*result)
}
