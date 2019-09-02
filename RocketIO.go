package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ReadCsvNumerical :
func ReadCsvNumerical(fname string, dropFrist bool) Matrix {
	var data Matrix

	var f, err = os.Open(fname)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	defer f.Close()
	var converter = func(s []string) []float64 {
		var temp = make([]float64, len(s), len(s))
		for index, value := range s {
			t, err := strconv.ParseFloat(value, 64)
			if err != nil {
				fmt.Println(err.Error())
				return nil
			}
			temp[index] = t
		}
		return temp
	}
	var scanner = bufio.NewScanner(f)
	if dropFrist {
		scanner.Scan()
		_ = scanner.Text()
	}
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), ",")
		list := converter(row)
		data = append(data, list)
	}
	return data
}

// ReadCsvComplete :
func ReadCsvComplete(fname string) (DataFrame, error) {
	var data [][]string
	var f, err = os.Open(fname)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		row := strings.Split(string(scanner.Text()), ",")
		data = append(data, row)
	}
	return data, nil
}

// WriteCsv :
func WriteCsv(m Matrix, fname string) (int, error) {
	f, err := os.Create(fname)
	defer f.Close()
	var str string
	var n int
	var writer = bufio.NewWriter(f)
	for _, value := range m {
		for col, i := range value {
			str = str + strconv.FormatFloat(i, 'f', 6, 64)
			if col != m.Col()-1 {
				str = str + ","
			}
		}
		str = str + "\n"
		n, err = writer.WriteString(str)
		writer.Flush()
		str = ""
	}
	return n, err
}
