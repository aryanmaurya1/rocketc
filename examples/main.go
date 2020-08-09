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

	var m rocketc.Matrix
	m = rocketc.Ones(500, 40)
	var f = func(v float32) float32 {
		return v*2 + 2
	}
	m.Map(f, true)
	fmt.Println(rocketc.VStack(m, m, m))
}
