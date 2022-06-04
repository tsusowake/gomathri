package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewFloat(0.1)
	b := big.NewFloat(0.2)
	c := new(big.Float).Add(a, b)

	// c, floating pointing error...
	fmt.Println(a, b, c)
}