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
}
