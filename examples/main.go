package main

import (
	"fmt"

	"github.com/aryanmaurya1/rocketc"
)

func main() {
	var d rocketc.DataFrame
	d, err := rocketc.ReadCSVDataFrame("./datasets/USA_Housing.csv")
	if err != nil {
		fmt.Println(err.Error())
	}
	rocketc.PrintDataframe(d.Head(10))

	var m = rocketc.Matrix{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {1, 2, 3}}

	rocketc.PrintMatrix(m)
	rocketc.PrintMatrix(rocketc.Sum(m, 1))
	rocketc.PrintMatrix(rocketc.Min(m, 1))
	rocketc.PrintMatrix(rocketc.Max(m, 1))
	rocketc.PrintMatrix(rocketc.Mean(m, 1))
	rocketc.PrintMatrix(rocketc.GetColumnsMatrix(m, 0, 2))

}
