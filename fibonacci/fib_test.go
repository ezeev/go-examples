package fibonacci_test

import (
	"testing"

	"github.com/ezeev/go-examples/fibonacci"
)

func TestFib(t *testing.T) {
	fibonacci.PrintFib(50)
}

func TestFibRecur(t *testing.T) {
	for i := 0; i <= 40; i++ {
		t.Logf(" %d", fibonacci.SlowFib(i))
	}
}

func TestFibMemo(t *testing.T) {

	fibonacci.FibMemoization(40)

}
