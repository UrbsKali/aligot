// test for utils.go

package utils

import (
	"testing"
)

func TestMatrixString(t *testing.T) {
	test := Matrix{{1, 2, 3}, {4, 5, 6}}
	expected := "[1 2 3]\n[4 5 6]\n"
	if test.String() != expected {
		t.Errorf("Matrix.String() = %v, want %v", test.String(), expected)
	}
}

func TestMatrixRows(t *testing.T) {
	test := Matrix{{1, 2, 3}, {4, 5, 6}}
	expected := 2
	if test.Rows() != expected {
		t.Errorf("Matrix.Rows() = %v, want %v", test.Rows(), expected)
	}
}

func TestMatrixCols(t *testing.T) {
	test := Matrix{{1, 2, 3}, {4, 5, 6}}
	expected := 3
	if test.Cols() != expected {
		t.Errorf("Matrix.Cols() = %v, want %v", test.Cols(), expected)
	}
}

func TestMatrixTranspose(t *testing.T) {
	test := Matrix{{1, 2, 3}, {4, 5, 6}}
	expected := Matrix{{1, 4}, {2, 5}, {3, 6}}
	if !MatrixEqual(test.Transpose(), expected) {
		t.Errorf("Matrix.Transpose() = %v, want %v", test.Transpose(), expected)
	}
}

func TestMatrixMultiply(t *testing.T) {
	test := Matrix{{1, 2, 3}, {4, 5, 6}}
	test2 := Matrix{{1, 2}, {3, 4}, {5, 6}}
	expected := Matrix{{22, 28}, {49, 64}}
	if !MatrixEqual(test.Multiply(&test2), expected) {
		t.Errorf("Matrix.Multiply() = %v, want %v", test.Multiply(&test2), expected)
	}
}

func TestMatrixEqual(t *testing.T) {
	test := Matrix{{1, 2, 3}, {4, 5, 6}}
	test2 := Matrix{{1, 2, 3}, {4, 5, 6}}
	if !MatrixEqual(test, test2) {
		t.Errorf("MatrixEqual() = %v, want %v", MatrixEqual(test, test2), true)
	}
}
