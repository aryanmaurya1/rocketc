package rocketc

// Matrix : Basic numerical container (2D array), core datatype.
type Matrix [][]float32

// Init : Initializes a Matrix to 1 X 1 Matrix containing zero.
func (m *Matrix) Init() {
	*m = make(Matrix, 1)
	(*m)[0] = make([]float32, 1)
}

// Rows : Returns number of rows in Matrix.
func (m Matrix) Rows() int {
	return len(m)
}

// Cols : Returns number of columns in Matrix.
func (m Matrix) Cols() int {
	if len(m) != 0 {
		return len(m[0])
	}
	return 0
}

// Shape : Returns a int array []int containing the dimensions of Matrix.
func (m Matrix) Shape() [2]int {
	var s [2]int
	s[0] = m.Rows()
	s[1] = m.Cols()
	return s
}

// MakeMatrixUniform : Makes Matrix rows of same size,
// by filling rest of the row with zero value of float32
// if different size rows are present in Matrix.
func (m *Matrix) MakeMatrixUniform() {
	nRows := m.Rows()
	sizeMaxRows := 0
	for i := 0; i < nRows; i++ {
		if len((*m)[i]) > sizeMaxRows {
			sizeMaxRows = len((*m)[i])
		}
	}
	for i := 0; i < nRows; i++ {
		if len((*m)[i]) < sizeMaxRows {
			temp := make([]float32, sizeMaxRows-len((*m)[i]))
			(*m)[i] = append((*m)[i], temp...)
		}
	}
}

// Fill : fills the Matrix with a given number i.
func (m *Matrix) Fill(i float32) {
	for row := range *m {
		for col := range (*m)[row] {
			(*m)[row][col] = i
		}
	}
}

// Transpose : Transpose the Matrix, takes a boolean inplace if True
// transposes Matrix inplace.
func (m *Matrix) Transpose(inplace bool) Matrix {
	var s = Zeros(m.Cols(), m.Rows())
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

// Ones : Fills the Matrix with all zeros.
func (m *Matrix) Ones() {
	(*m).Fill(1)
}

// Zeros : Fills the Matrix with all ones.
func (m *Matrix) Zeros() {
	(*m).Fill(0)
}

// Add : Adds a given number to every element of Matrix, takes a boolean inplace if True
// performs addition inplace.
func (m *Matrix) Add(i float32, inplace bool) Matrix {
	var s = Zeros(m.Rows(), m.Cols())
	s.Fill(i)
	s = AddElementwise(s, *m)
	if inplace {
		*m = s
	}
	return s
}

// Sub : Subtracts a given number from every element of Matrix, takes a boolean inplace if True
// performs subtraction inplace.
func (m *Matrix) Sub(i float32, inplace bool) Matrix {
	var s = Zeros((*m).Rows(), (*m).Cols())
	s.Fill(i)
	s = AddElementwise(s, *m)
	if inplace {
		*m = s
	}
	return s
}

// Mul : Multiplies i to every element of the Matrix, takes a boolean inplace if True
// performs multiplication inplace.
func (m *Matrix) Mul(i float32, inplace bool) Matrix {
	var s = Zeros((*m).Rows(), (*m).Cols())
	s.Fill(i)
	s = MulElementwise(*m, s)
	if inplace {
		*m = s
	}
	return s
}

// Div : Divides every element of Matrix by i, takes a boolean inplace if True
// performs division inplace.
func (m *Matrix) Div(i float32, inplace bool) Matrix {
	var s = Zeros((*m).Rows(), (*m).Cols())
	s.Fill(i)
	s = DivElementwise(*m, s)
	if inplace {
		*m = s
	}
	return s
}

// ReciproElementwise : Elementwise reciprocal of Matrix elements, takes a boolean inplace if True
// performs reciprocal inplace.
func (m *Matrix) ReciproElementwise(inplace bool) Matrix {
	var s = Zeros((*m).Rows(), (*m).Cols())
	for row := range *m {
		for col := range (*m)[row] {
			s[row][col] = 1 / (*m)[row][col]
		}
	}
	if inplace {
		*m = s
	}
	return s
}

// Map : Applies given function to every element of Matrix,
// takes a boolean inplace if True performs mapping inplace.
func (m *Matrix) Map(f func(v float32) float32, inplace bool) Matrix {
	var temp = Zeros(m.Rows(), m.Cols())
	for i := 0; i < len(temp); i++ {
		for j := 0; j < len(temp[i]); j++ {
			temp[i][j] = f((*m)[i][j])
		}
	}
	if inplace {
		*m = temp
	}
	return temp
}
