package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Matrix : Basic numerical container (2D array), core datatype
type Matrix [][]float32

// Init : Initializes a matrix to 1 X 1 matrix containing zero
func (m *Matrix) Init() {
	*m = make(Matrix, 1, 1)
	(*m)[0] = make([]float32, 1, 1)
}

// Row : Returns number of rows in matrix (returned value -1 represent nil matrix)
func (m *Matrix) Row() int {
	return len(*m)
}

// Col : Returns number of columns in matrix (returned value -1 represent nil matrix)
func (m *Matrix) Col() int {
	if len(*m) != 0 {
		return len((*m)[0])
	}
	return 0
}

// Shape : Returns a float32 slice containing the dimensions of matrix
func (m *Matrix) Shape() []int {
	var s = make([]int, 2, 2)
	s[0] = (*m).Row()
	s[1] = (*m).Col()
	return s
}

// Fill : fills the matrix with a given number
func (m *Matrix) Fill(i float32) {
	for row := range *m {
		for col := range (*m)[row] {
			(*m)[row][col] = i
		}
	}
}

// Transpose : transpose the matrix
func (m *Matrix) Transpose(inplace bool) Matrix {
	var s = Zeros((*m).Col(), (*m).Row())
	for row := range *m {
		for col := range (*m)[row] {
			s[col][row] = (*m)[row][col]
		}
	}
	if inplace {
		*m = s
	}
	return s
}

// Ones : Fills the matrix with all zeros
func (m *Matrix) Ones() {
	(*m).Fill(1)
}

// Zeros : Fills the matrix with all ones
func (m *Matrix) Zeros() {
	(*m).Fill(0)
}

// Add : Adds a given number to every element of matrix
func (m *Matrix) Add(i float32, f bool) Matrix {
	var s = Zeros((*m).Row(), (*m).Col())
	s.Fill(i)
	s = AddElementwise(s, *m)
	if f {
		*m = s
	}
	return s
}

// Sub : Subtracts a given number from every element of matrix
func (m *Matrix) Sub(i float32, f bool) Matrix {
	var s = Zeros((*m).Row(), (*m).Col())
	s.Fill(i)
	s = AddElementwise(s, *m)
	if f {
		*m = s
	}
	return s
}

// Mul : Multiplies i to every element of the matrix
func (m *Matrix) Mul(i float32, f bool) Matrix {
	var s = Zeros((*m).Row(), (*m).Col())
	s.Fill(i)
	s = MulElementwise(*m, s)
	if f {
		*m = s
	}
	return s
}

// Div : performs elementwise division
func (m *Matrix) Div(i float32, f bool) Matrix {
	var s = Zeros((*m).Row(), (*m).Col())
	s.Fill(i)
	s = DivElementwise(*m, s)
	if f {
		*m = s
	}
	return s
}

// ReciproElementwise : elementwise recipro of matrix elements
func (m *Matrix) ReciproElementwise(f bool) Matrix {
	var s = Zeros((*m).Row(), (*m).Col())
	for row := range *m {
		for col := range (*m)[row] {
			s[row][col] = 1 / (*m)[row][col]
		}
	}
	if f {
		*m = s
	}
	return s
}

// PrintMatrix : prints matrix
func PrintMatrix(m ...*Matrix) {
	lambda := func(m *Matrix) {
		for i := 0; i < len((*m)[0]); i++ {
			fmt.Printf("%10d ", i)
		}
		fmt.Println()
		for i := 0; i < len((*m)[0])*6; i++ {
			fmt.Printf("--")
		}
		fmt.Println()
		for row := range *m {
			fmt.Printf("%3d |", row)
			for col := range (*m)[row] {
				fmt.Printf("%10.6f ", (*m)[row][col])
			}
			// fmt.Printf(",")
			fmt.Println()
			// fmt.Println((*m)[row])
		}
	}
	for _, value := range m {
		lambda(value)
		fmt.Println()
	}
}

// DimensionEqual : checks if two matrix has same dimension
func DimensionEqual(m1, m2 Matrix) bool {
	if m1.Row() == m2.Row() && m1.Col() == m2.Col() {
		return true
	}
	return false
}

// AddElementwise : adds two matrix
func AddElementwise(m1, m2 Matrix) Matrix {
	if !DimensionEqual(m1, m2) {
		fmt.Printf("Dimensions are not equal !!!!")
		return nil
	}
	var s = Zeros(m1.Row(), m1.Col())
	for row := range s {
		for col := range s[row] {
			s[row][col] = m1[row][col] + m2[row][col]
		}
	}
	return s
}

// SubElementwise : subtracts two matrices
func SubElementwise(m1, m2 Matrix) Matrix {
	if !DimensionEqual(m1, m2) {
		fmt.Printf("Dimensions are not equal !!!!")
		return nil
	}
	var s = Zeros(m1.Row(), m1.Col())
	for row := range s {
		for col := range s[row] {
			s[row][col] = m1[row][col] - m2[row][col]
		}
	}
	return s
}

// MulElementwise : Elementwise multiplication of elements in matrx m1 and m2.
func MulElementwise(m1, m2 Matrix) Matrix {
	if !DimensionEqual(m1, m2) {
		fmt.Println("Dimension Mismatch !!!")
		return nil
	}
	var s = Zeros(m1.Row(), m2.Col())
	for row := 0; row < m1.Row(); row++ {
		for col := 0; col < m2.Col(); col++ {
			s[row][col] = m2[row][col] * m1[row][col]
		}
	}
	return s
}

// DivElementwise : Elementwise division
func DivElementwise(m1, m2 Matrix) Matrix {
	var m = m2
	m.ReciproElementwise(true)
	return MulElementwise(m1, m)
}

// Dot :
func Dot(m1, m2 Matrix) Matrix {
	isPossible := func(m1, m2 Matrix) bool {
		if m1.Col() == m2.Row() {
			return true
		}
		return false
	}
	if !isPossible(m1, m2) {
		fmt.Println("Dimension Mismatch !!!")
		return nil
	}
	var s = Zeros(m1.Row(), m2.Col())
	for row, value := range s {
		var sum float32
		for col := range value {
			for i := 0; i < m1.Col(); i++ {
				sum += m1[row][i] * m2[i][col]
			}
			s[row][col] = sum
			sum = 0
		}
	}
	return s
}

// Zeros : Matrix of zeros
func Zeros(row, col int) Matrix {
	var m = make(Matrix, row, row)
	for index := range m {
		m[index] = make([]float32, col, col)
	}
	return m
}

// Ones : Matrix of Ones
func Ones(row, col int) Matrix {
	var m = Zeros(row, col)
	m.Fill(1)
	return m
}

// Sum : performs the sum of elements in matrix
func Sum(m Matrix, axis int) Matrix {
	if axis > 1 || axis < 0 {
		fmt.Println("Error in Axis")
		return nil
	}
	if axis == 0 {
		var s = Zeros(1, m.Col())
		for i := 0; i < len(s[0]); i++ {
			for _, value := range m {
				s[0][i] += value[i]
			}
		}
		return s
	}
	var s = Zeros(m.Row(), 1)
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

// Random :
func Random(row, col int, n float32) Matrix {
	if row < 0 || col < 0 {
		fmt.Println("Wrong Dimensions !!!")
		return nil
	}
	rand.Seed(time.Now().UnixNano())
	var s = Zeros(row, col)
	for row, value := range s {
		for col := range value {
			s[row][col] = rand.Float32() * n
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
	var s = make(Matrix, m[0].Row(), m[0].Row())
	for _, value := range m {
		for index, row := range value {
			s[index] = append(s[index], row...)
		}
	}
	return s
}

// Max :
func Max(m Matrix, axis int) Matrix {
	if axis == 0 {
		var s = Zeros(1, m.Col())
		for row, value := range m {
			for col := range value {
				if m[row][col] > s[0][col] {
					s[0][col] = m[row][col]
				}
			}
		}
		return s
	}
	var s = Zeros(m.Row(), 1)
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
		var s = Zeros(1, m.Col())
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
	var s = Zeros(m.Row(), 1)
	for row, value := range m {
		s[row][0] = minimum(value)
	}
	return s
}

// Mean :
func Mean(m Matrix, axis int) Matrix {
	if axis == 0 {
		var l = m.Row()
		var s = Sum(m, axis)
		for col := range s[0] {
			s[0][col] /= float32(l)
		}
		return s
	}
	var l = m.Col()
	var s = Sum(m, 1)
	for col := range s {
		s[col][0] /= float32(l)
	}
	return s
}

// GetColumnsMatrix :
func GetColumnsMatrix(m Matrix, i ...int) Matrix {
	var c = Zeros(m.Row(), len(i))
	for row, value := range m {
		for index, v := range i {
			c[row][index] = value[v]
		}
	}
	return c
}
