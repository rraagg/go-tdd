package calculator_test

import (
	"testing"

	"github.com/rraagg/go-tdd/calculator"
)

func TestAdd(t *testing.T) {
	// Arrange
	e := calculator.Engine{}
	x, y := 2.5, 3.5
	want := 6.0
	// Act
	got := e.Add(x, y)
	// Asser
	if got != 6.0 {
		t.Errorf(
			"Add(%.2f,%.2f) incorrect, got: %.2f, want: %.2f",
			x,
			y,
			got,
			want,
		)
	}
}
