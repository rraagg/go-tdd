package calculator_test

import (
	"fmt"
	"math"
	"os"
	"testing"

	"github.com/rraagg/go-tdd/calculator"
)

func init() {
	fmt.Println("init setup")
}

func setup() {
	fmt.Println("setting up")
}

func teardown() {
	fmt.Println("tearing down")
}

func TestMain(m *testing.M) {
	// setup statements
	setup()
	// run the tests
	e := m.Run()
	// cleanup statements
	teardown()
	// report the exit code
	os.Exit(e)
}

func BenchmarkAdd(b *testing.B) {
	e := calculator.Engine{}
	// run the add function b.N times
	for i := 0; i < b.N; i++ {
		e.Add(2, 3)
	}
}

func TestAdd(t *testing.T) {
	// Arrange
	defer func() {
		fmt.Println("tearing down TestAdd")
	}()
	e := calculator.Engine{}
	x, y := 2.5, 3.5
	want := 6.0
	// Act
	got := e.Add(x, y)
	// Assert
	if got != want {
		t.Errorf(
			"Add(%.2f,%.2f) incorrect, got: %.2f, want: %.2f",
			x,
			y,
			got,
			want,
		)
	}
}

func TestSubtract(t *testing.T) {
	// Arrange
	e := calculator.Engine{}
	x, y := 5.0, 3.5
	want := 1.5
	// Act
	got := e.Subtract(x, y)
	// Assert
	if got != want {
		t.Errorf(
			"Subtract(%.2f, %.2f) incorrect, got: %.2f, want: %.2f",
			x,
			y,
			got,
			want,
		)
	}
}

func TestMultiply(t *testing.T) {
	// Arrange
	e := calculator.Engine{}
	x, y := 6.0, 3.5
	want := 21.0
	// Act
	got := e.Multiply(x, y)
	// Assert
	if got != want {
		t.Errorf(
			"Multiply(%.2f, %.2f) incorrect, got: %.2f, want: %.2f",
			x,
			y,
			got,
			want,
		)
	}
}

func TestDivide(t *testing.T) {
	// Arrange
	e := calculator.Engine{}
	x, y := 9.9, 3.3
	want := 3.0
	// Act
	got := e.Divide(x, y)
	// Assert
	if math.Round(got) != want {
		t.Errorf(
			"Divide(%.2f, %.2f) incorrect, got: %.2f, want %.2f",
			x,
			y,
			got,
			want,
		)
	}
}
