package main

import (
	"fmt"
	"math"
)


//Go's if statements are like its for loops; the expression need not be surrounded by parentheses ( ) but the braces { } are required.

//Go's if statements can include an optional statement (v := math.Pow(x, n)) that is executed before the condition (v < limit) is evaluated.


func pow(x, n, limit float64) float64 {
	if v := math.Pow(x, n); v < limit {
		return v  //either returns v or limit
	}
	return limit //return v will throw an error here: because v is declared within the 'if' scope

}

func main() {
	fmt.Println(
		pow(3, 2, 10), // This will return 9 because 3^2 = 9; 9 < 10 = lim
		pow(3, 3, 20), // This will return 20 because 3^3 = 27; 27 > 20 = lim
		pow(3, 4, 28), // This will return 28 because 3^4 = 81; 81 > 28 = lim
	)
}
