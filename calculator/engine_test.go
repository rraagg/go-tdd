package calculator_test

import (
	"testing"

	"github.com/rraagg/go-tdd/calculator"
)

func TestAdd(t *testing.T) {
	// Arrange
	e := calculator.Engine{}
	// Act
	got := e.Add(2.5, 3.5)
	// Asser
	if got != 6.0 {
		t.Errorf(
			"Add(%.2f,%.2f) incorrect, got: %.2f, want: %.2f",
			2.5,
			3.5,
			got,
			6.0,
		)
	}
}
