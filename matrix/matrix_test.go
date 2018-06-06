package matrix_test

import (
	"testing"

	"github.com/ezeev/go-examples/matrix"
)

func TestMatrix(t *testing.T) {

	mat := matrix.Matrix(2)
	t.Log(mat)

	mat = matrix.Matrix(3)
	t.Log(mat)

	mat = matrix.SpiralMatrix(3)
	t.Log(mat)

	mat = matrix.SpiralMatrix(5)
	t.Log(mat)

}
