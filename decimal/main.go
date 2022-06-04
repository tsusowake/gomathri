package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {
	fmt.Println("decimal")

	f1 := decimal.NewFromFloat(10000000000000000.0)
	fmt.Println(f1.String())

	d1 := f1.Div(decimal.NewFromFloat(3))
	fmt.Println(d1) // out: 3333333333333333.3333333333333333

	f2, e := d1.Float64()
	fmt.Println(f2)
	fmt.Println(e)
	fmt.Println(d1.String())
}
