package rocketc

import (
	"fmt"
	"math/rand"
	"time"
)

// PrintMatrix : Prints Matrix in a pretty manner.
func PrintMatrix(m ...Matrix) {
	lambda := func(m Matrix) {
		for i := 0; i < len(m[0]); i++ {
			fmt.Printf("%10d ", i)
		}
		fmt.Println()
		for i := 0; i < len(m[0])*6; i++ {
			fmt.Printf("--")
		}
		fmt.Println()
		for row := range m {
			fmt.Printf("%3d |", row)
			for col := range m[row] {
				fmt.Printf("%10.6f ", m[row][col])
			}
			fmt.Println()
		}
	}
	for _, value := range m {
		lambda(value)
		fmt.Println()
	}
}

// CopyMatrix : Returns the copy of src Matrix.
func CopyMatrix(src Matrix) Matrix {
	m := Zeros(src.Rows(), src.Cols())
	copy(m, src)
	return m
}

// DimensionEqual : Checks if m1 and m2 has same dimension, returns boolean.
func DimensionEqual(m1, m2 Matrix) bool {
	if (m1.Rows() == m2.Rows()) && (m1.Cols() == m2.Cols()) {
		return true
	}
	return false
}

// AddElementwise : Adds two Matrix m1 and m2.
func AddElementwise(m1, m2 Matrix) Matrix {
	if !DimensionEqual(m1, m2) {
		panic("Dimensions are not equal.")
	}
	var s = Zeros(m1.Rows(), m1.Cols())
	for row := range s {
		for col := range s[row] {
			s[row][col] = m1[row][col] + m2[row][col]
		}
	}
	return s
}

// SubElementwise : Subtracts two Matrix m1 and m2.
func SubElementwise(m1, m2 Matrix) Matrix {
	if !DimensionEqual(m1, m2) {
		panic("Dimensions are not equal.")
	}
	var s = Zeros(m1.Rows(), m1.Cols())
	for row := range s {
		for col := range s[row] {
			s[row][col] = m1[row][col] - m2[row][col]
		}
	}
	return s
}

// MulElementwise : Elementwise multiplication of Matrix m1 and m2.
func MulElementwise(m1, m2 Matrix) Matrix {
	if !DimensionEqual(m1, m2) {
		panic("Dimensions are not equal.")
	}
	var s = Zeros(m1.Rows(), m2.Cols())
	for row := 0; row < m1.Rows(); row++ {
		for col := 0; col < m2.Cols(); col++ {
			s[row][col] = m2[row][col] * m1[row][col]
		}
	}
	return s
}

// DivElementwise : Elementwise division of Matric m1 and m2.
func DivElementwise(m1, m2 Matrix) Matrix {
	var m = m2.ReciproElementwise(false)
	return MulElementwise(m1, m)
}

// Multiply : Returns the result of Matrix Multiplication of m1 and m2.
func Multiply(m1, m2 Matrix) Matrix {
	isPossible := func(m1, m2 Matrix) bool {
		if m1.Cols() == m2.Rows() {
			return true
		}
		return false
	}
	if !isPossible(m1, m2) {
		panic("Dimension Mismatch for matrix multiplication.")
	}
	var s = Zeros(m1.Rows(), m2.Cols())
	for row, value := range s {
		var sum float32
		for col := range value {
			for i := 0; i < m1.Cols(); i++ {
				sum += m1[row][i] * m2[i][col]
			}
			s[row][col] = sum
			sum = 0
		}
	}
	return s
}

// Zeros : Returns a Matrix of zeros of given rows and cols.
func Zeros(rows, cols int) Matrix {
	var m = make(Matrix, rows)
	for index := range m {
		m[index] = make([]float32, cols)
	}
	return m
}

// Ones : Returns a Matrix of ones of given rows and cols.
func Ones(rows, cols int) Matrix {
	var m = Zeros(rows, cols)
	m.Fill(1)
	return m
}

// Sum : Returns the sum of elements of Matrix according to given axis.
func Sum(m Matrix, axis int) Matrix {
	if axis > 1 || axis < 0 {
		panic("Error in Axis")
	}
	if axis == 0 {
		var s = Zeros(1, m.Cols())
		for i := 0; i < len(s[0]); i++ {
			for _, value := range m {
				s[0][i] += value[i]
			}
		}
		return s
	}
	var s = Zeros(m.Rows(), 1)
	lambda := func(l []float32) float32 {
		var s float32
		for _, value := range l {
			s += value
		}
		return s
	}
	for index, value := range m {
		s[index][0] = lambda(value)
	}
	return s
}

// Random : Generates and returns a Matrix of given rows and cols
// containing random numbers.
func Random(rows, cols int) Matrix {
	if (rows <= 0) || (cols <= 0) {
		return nil
	}
	rand.Seed(time.Now().UnixNano())
	var s = Zeros(rows, cols)
	for row, value := range s {
		for col := range value {
			s[row][col] = rand.Float32()
		}
	}
	return s
}

// VStack :
func VStack(m ...Matrix) Matrix {
	var s Matrix
	for _, value := range m {
		for _, row := range value {
			s = append(s, row)
		}
	}
	return s
}

// HStack :
func HStack(m ...Matrix) Matrix {
	var s = make(Matrix, m[0].Rows())
	for _, value := range m {
		for index, row := range value {
			s[index] = append(s[index], row...)
		}
	}
	return s
}

// Max :
// 0 -> Along the rows
// 1 -> Along the columns
func Max(m Matrix, axis int) Matrix {
	if axis == 1 {
		var s = Zeros(1, m.Cols())
		for row, value := range m {
			for col := range value {
				if m[row][col] > s[0][col] {
					s[0][col] = m[row][col]
				}
			}
		}
		return s
	}
	var s = Zeros(m.Rows(), 1)
	for row, value := range m {
		for col := range value {
			if m[row][col] > s[row][0] {
				s[row][0] = m[row][col]
			}
		}
	}
	return s
}

// Min :
func Min(m Matrix, axis int) Matrix {
	minimum := func(m []float32) float32 {
		min := m[0]
		for _, value := range m {
			if value < min {
				min = value
			}
		}
		return min
	}
	if axis == 0 {
		var s = Zeros(1, m.Cols())
		s[0] = m[0]
		for row, value := range m {
			for col := range value {
				if m[row][col] < s[0][col] {
					s[0][col] = m[row][col]
				}
			}
		}
		return s
	}
	var s = Zeros(m.Rows(), 1)
	for row, value := range m {
		s[row][0] = minimum(value)
	}
	return s
}

// Mean :
func Mean(m Matrix, axis int) Matrix {
	if axis == 0 {
		var l = m.Rows()
		var s = Sum(m, axis)
		for col := range s[0] {
			s[0][col] /= float32(l)
		}
		return s
	}
	var l = m.Cols()
	var s = Sum(m, 1)
	for col := range s {
		s[col][0] /= float32(l)
	}
	return s
}

// GetColumnsMatrix :
func GetColumnsMatrix(m Matrix, i ...int) Matrix {
	var c = Zeros(m.Rows(), len(i))
	for row, value := range m {
		for index, v := range i {
			c[row][index] = value[v]
		}
	}
	return c
}
