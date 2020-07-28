package main

import (
	"fmt"
	"strconv"
)

// DataFrame : Basic data container
type DataFrame [][]string

// Row : Returns number of rows in DataFrame
func (d *DataFrame) Row() int {
	return len(*d)
}

// Col : Returns number of columns in DataFrame
func (d *DataFrame) Col() int {
	if len(*d) != 0 {
		return len((*d)[0])
	}
	return 0
}

// Shape : Returns shape of DataFrame
func (d *DataFrame) Shape() []int {
	var size = make([]int, 2, 2)
	size[0] = (*d).Row()
	size[1] = (*d).Col()
	return size
}

// Headers : Returns header of the dataframe i.e row 0
func (d *DataFrame) Headers() []string {
	return (*d)[0]
}

// Head : Returns first n rows of DataFrame
func (d *DataFrame) Head(n int64) DataFrame {
	return (*d)[0:n]
}

// Allocate : Allocate a blank DataFrame of given size
func Allocate(row, col int) DataFrame {
	var d = make(DataFrame, row, row)
	for i := 0; i < row; i++ {
		d[i] = make([]string, col, col)
	}
	return d
}

// WipeDown : returns unifom DataFrame by only including rows of length l
// in returned DataFrame
func WipeDown(m DataFrame, l int) DataFrame {
	var r DataFrame
	n := m.Row()
	for i := 0; i < n; i++ {
		value := m[i]
		if len(value) == l {
			r = append(r, value)
		}
	}
	return r
}

// DropColumn :
func DropColumn(d DataFrame, i ...int) DataFrame {
	f := func(arr []int) int {
		var max = arr[0]
		for _, value := range arr {
			if value > max {
				max = value
			}
		}
		return max
	}
	var result = make(DataFrame, len(d), len(d))
	var arr = make([]int, f(i)+1, f(i)+1)
	for _, value := range i {
		arr[value]++
	}
	for j := 0; j < len(d[0]); j++ {
		for i := 0; i < len(d); i++ {
			if arr[j] > 0 {
				break
			}
			result[i] = append(result[i], d[i][j])
		}
	}
	return result
}

// ConvMatrix :
func ConvMatrix(d DataFrame) Matrix {
	var m = Zeros(d.Row(), d.Col())
	var r = d.Row()
	var c = d.Col()
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			temp, err := strconv.ParseFloat(d[i][j], 64)
			if err != nil {
				fmt.Println(err.Error())
			}
			m[i][j] = float32(temp)
		}
	}
	return m
}

// PrintDataframe :
func PrintDataframe(m ...*DataFrame) {
	lambda := func(m *DataFrame) {
		// for i := 0; i < (*m).Col(); i++ {
		// 	fmt.Printf("%15d ", i)
		// }
		fmt.Println()
		for i := 0; i < len((*m)[0])*9; i++ {
			fmt.Printf("--")
		}
		fmt.Println()
		for row := range *m {
			fmt.Printf("%3d |", row)
			for col := range (*m)[row] {
				fmt.Printf("%-16s ", (*m)[row][col])
			}
			fmt.Println()
		}
	}
	for _, value := range m {
		lambda(value)
		fmt.Println()
	}
}

// GetColumnsDataFrame :
func GetColumnsDataFrame(d DataFrame, i ...int) DataFrame {
	var c = Allocate(d.Row(), len(i))
	for row, value := range d {
		for index, v := range i {
			c[row][index] = value[v]
		}
	}
	return c
}
