package math

import (
	"fmt"
	"log"
)

// A + B
func Sum(A, B Matrix) Matrix {
	c := make(Matrix)
	rowsA := len(A)
	colsB := getCols(B)
	for i := 0; i < rowsA; i++ {
		for j := 0; j < colsB; j++ {
			c.Add(i, j, A[i][j]+B[i][j])
		}
	}
	return c
}

// A * B
func Multiplay(A, B Matrix) Matrix {
	c := make(Matrix)
	rowsA := len(A)
	rowsB := len(B)
	colsA := getCols(A)
	colsB := getCols(B)

	if colsA != rowsB {
		fmt.Println("The number of columns in A must equal the number of rows in B")
		log.Fatalf("A[_,%d] B[%d,_]\n", colsA, rowsB)
		return c
	}

	for i := 0; i < rowsA; i++ {
		for j := 0; j < colsB; j++ {
			for k := 0; k < colsA; k++ {
				c.Add(i, j, A[i][k]*B[k][j])
			}
		}
	}
	return c
}

// Apply dsigmoid function on all matrix elements
func Dsigmoid(A Matrix) Matrix {
	c := make(Matrix)
	rowsA := len(A)
	colsA := getCols(A)
	for i := 0; i < rowsA; i++ {
		for j := 0; j < colsA; j++ {
			c.Set(i, j, A[i][j]*(1-A[i][j]))
		}
	}
	return c
}

// Transpose
func Transpose(A Matrix) Matrix {
	c := make(Matrix)
	rowsA := len(A)
	colsA := getCols(A)
	for i := 0; i < colsA; i++ {
		for j := 0; j < rowsA; j++ {
			c.Set(i, j, A[j][i])
		}
	}
	return c
}

// A ∙ B
func Hadamard(A, B Matrix) Matrix {
	c := make(Matrix)
	rowsA := len(A)
	colsA := getCols(A)
	for i := 0; i < rowsA; i++ {
		for j := 0; j < colsA; j++ {
			c.Set(i, j, A[i][j]*B[i][j])
		}
	}
	return c
}

// A × B
func ScalarMultiply(A Matrix, scalar float64) Matrix {
	c := make(Matrix)
	rowsA := len(A)
	colsA := getCols(A)
	for i := 0; i < rowsA; i++ {
		for j := 0; j < colsA; j++ {
			c.Set(i, j, A[i][j]*scalar)
		}
	}
	return c
}
