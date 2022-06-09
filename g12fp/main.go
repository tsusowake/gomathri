package main

import (
	"fmt"
	"strconv"
)

var values = map[int64]string{
	1: "10000.123456789012345",
	2: "10000.123456789012345",
	//	3: "10000.123456789012345",
	//	4: "10000.123456789012345",
}

func main() {
	var total float64
	for _, fs := range values {
		fv, err := strconv.ParseFloat(fs, 64)
		if err != nil {
			panic(err)
		}
		fmt.Println("value: ", fmt.Sprintf("%g", fv))
		total += fv
	}

	fmt.Println("total: ", fmt.Sprintf("%g", total))
}
