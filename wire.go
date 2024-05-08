//go:build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/rraagg/go-tdd/calculator"
)

var Set = wire.NewSet(calculator.NewEngine,
	wire.Bind(new(calculator.Adder),
		new(*calculator.Engine)), calculator.NewCalculator)

func InitCalc() *calculator.Calculator {
	wire.Build(Set)
	return nil
}
