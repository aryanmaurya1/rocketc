// Package rocketc provides basic functionalities to work with CSV data
// and 2D Matrices.
package rocketc

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ReadCSVMatrix : Use this function to read CSV if you are sure that CSV
// file only contains parsable numerical values (float64).Takes a string filename
// and a boolean whether to drop first row or not. Returns a Matrix.
// Note : Drop first row if it contains name of columns.
func ReadCSVMatrix(fname string, dropFirst bool) (Matrix, error) {
	var data Matrix

	var f, err = os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var converter = func(s []string) []float32 {
		var temp = make([]float32, len(s))
		for index, value := range s {
			t, err := strconv.ParseFloat(value, 64)
			if err != nil {
				fmt.Println(err.Error())
				return nil
			}
			temp[index] = float32(t)
		}
		return temp
	}
	var scanner = bufio.NewScanner(f)
	if dropFirst {
		scanner.Scan()
		_ = scanner.Text()
	}
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), ",")
		list := converter(row)
		data = append(data, list)
	}
	return data, nil
}

// ReadCSVDataFrame : Use this function to read CSV file completely.
// Takes a string filename. Returns DataFrame and not nil error value in case
// of any error occured. Currently ReadCSVDataFrame can only read matrices
// which do not have multiple values in single column.
// Note : Use this function if data contains both numeric and string values.
func ReadCSVDataFrame(fname string) (DataFrame, error) {
	var f, err = os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	reader := csv.NewReader(f)
	return reader.ReadAll()
}

// WriteCSVMatrix : Writes Matrix into a file, so that Matrix can be
// saved to disk for furthur. Takes a Matrix, slice of strings which are
// headers and filename as arguments. Returns an error value in case of
// any error occured.
// Note : This function can only write Matrix to file.
func WriteCSVMatrix(m Matrix, fname string) error {
	f, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer f.Close()
	writerF := func(value []float32, writer *bufio.Writer) error {
		var str string
		for col, i := range value {
			str = str + strconv.FormatFloat(float64(i), 'f', 6, 64)
			if col != m.Cols()-1 {
				str = str + ","
			}
		}
		str = str + "\n"
		_, err = writer.WriteString(str)
		if err != nil {
			return err
		}
		writer.Flush()
		str = ""
		return nil
	}
	var writer = bufio.NewWriter(f)
	for _, value := range m {
		err = writerF(value, writer) // start concatenation in empty string.
		if err != nil {
			return err
		}
	}
	return err
}

// WriteCSVDataFrame : Writes DataFrame into a file, so that DataFrame can be
// saved to disk for furthur. Takes a DataFrame and filename as
// arguments. Returns an error value in case of any error occured.
// Note : This function can only write DataFrame to file.
func WriteCSVDataFrame(d DataFrame, fname string) error {
	f, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer f.Close()
	var str string
	var writer = bufio.NewWriter(f)
	for _, value := range d {
		for col, i := range value {
			str = str + i
			if col != d.Cols()-1 {
				str = str + ","
			}
		}
		str = str + "\n"
		_, err = writer.WriteString(str)
		if err != nil {
			return err
		}
		writer.Flush()
		str = ""
	}
	return err
}
