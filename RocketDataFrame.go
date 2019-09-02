package main

import (
	"fmt"
	"strconv"
)

// DataFrame :
type DataFrame [][]string

// Row :
func (d *DataFrame) Row() int {
	return len(*d)
}

// Col :
func (d *DataFrame) Col() int {
	if len(*d) != 0 {
		return len((*d)[0])
	}
	return 0
}

// Shape :
func (d *DataFrame) Shape() []int {
	var size = make([]int, 2, 2)
	size[0] = (*d).Row()
	size[1] = (*d).Col()
	return size
}

// Head :
func (d *DataFrame) Head() {
	for index, value := range (*d)[0] {
		fmt.Println("Col :", index, " ", value)
	}
}

// PrintHead :
func (d *DataFrame) PrintHead(n int64) {
	var p = (*d)[0:n]
	PrintDataframe(&p)
}

// Allocate :
func Allocate(row, col int) DataFrame {
	var d = make(DataFrame, row, row)
	for index := range d {
		d[index] = make([]string, col, col)
	}
	return d
}

// WipeDown :
func WipeDown(m DataFrame, l int) DataFrame {
	var r DataFrame
	for _, value := range m {
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
			m[i][j] = temp
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
