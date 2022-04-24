package math

import (
	"fmt"
	"log"
	"math"
)

type Matrix (map[int]map[int]float64)

// get the size of the largest column of all the rows
func getCols(m Matrix) int {
	cols := 0
	for _, v := range m {
		if cols < len(v) {
			cols = len(v)
		}
	}
	return cols
}

// Apply sigmoid function on all matrix elements
func (m *Matrix) Sigmoid() {
	rows := len(*m)
	for i := 0; i < rows; i++ {
		cols := len((*m)[i])
		for j := 0; j < cols; j++ {
			m.Set(i, j, 1/(1+math.Exp(-((*m)[i][j]))))
		}
	}
}

// Set value on arbitrary position
func (m Matrix) Set(row, col int, v float64) {
	if m[row] == nil {
		m[row] = make(map[int]float64)
	}
	m[row][col] = v
}

// Add value on arbitrary position
func (m Matrix) Add(row, col int, v float64) {
	if m[row] == nil {
		m[row] = make(map[int]float64)
	}
	m[row][col] += v
}

// A + B
func (m Matrix) Sum(B Matrix) Matrix {
	c := make(Matrix)
	rowsA := len(m)
	colsB := getCols(B)
	for i := 0; i < rowsA; i++ {
		for j := 0; j < colsB; j++ {
			c.Add(i, j, m[i][j]+B[i][j])
		}
	}
	return c
}

// Subtract two values on arbitrary position
func (m Matrix) Diff(row, col int, v float64) {
	if m[row] == nil {
		m[row] = make(map[int]float64)
	}
	m[row][col] = v
}

// A - B
func (m Matrix) Sub(B Matrix) Matrix {
	c := make(Matrix)
	rowsA := len(m)
	colsB := getCols(B)
	for i := 0; i < rowsA; i++ {
		for j := 0; j < colsB; j++ {
			c.Diff(i, j, m[i][j]-B[i][j])
		}
	}
	return c
}

// A * B
func (m Matrix) Multiplay(B Matrix) Matrix {
	c := make(Matrix)
	rowsA := len(m)
	colsA := getCols(m)
	rowsB := len(B)
	colsB := getCols(B)

	if colsA != rowsB {
		fmt.Println("The number of columns in A must equal the number of rows in B")
		log.Fatalf("A[_,%d] B[%d,_]\n", colsA, rowsB)
		return c
	}

	for i := 0; i < rowsA; i++ {
		for j := 0; j < colsB; j++ {
			for k := 0; k < colsA; k++ {
				c.Add(i, j, m[i][k]*B[k][j])
			}
		}
	}
	return c
}
