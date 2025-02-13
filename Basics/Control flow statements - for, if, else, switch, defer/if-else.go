package main

import (
	"fmt"
	"math"
)

func pow(x, n, limit float64) float64 {
	if v := math.Pow(x, n); v < limit {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, limit)
	}
	return limit

	//if a function returns a value, it does not execute the code after the return statement
	//if 'return v' is executed, the code after it (along with 'return limit') will not be executed
}

func main() {
	fmt.Println(
		pow(3, 4, 28),
		pow(3, 2, 16),
		pow(3, 3, 20),
	)
}
