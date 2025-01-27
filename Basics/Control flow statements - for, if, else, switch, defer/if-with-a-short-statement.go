package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim //return v will throw an error here: because v is declared within the 'if' scope

}

func main() {
	fmt.Println(
		pow(3, 2, 10), // This will return 9 because 3^2 = 9; 9 < 10 = lim
		pow(3, 3, 20), // This will return 20 because 3^3 = 27; 27 > 20 = lim
		pow(3, 4, 28), // This will return 28 because 3^4 = 81; 81 > 28 = lim
	)
}
