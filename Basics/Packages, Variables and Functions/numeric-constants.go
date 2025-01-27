package main

import "fmt"

const (
	Big   = 1 << 100
	// Shift it left 100 places, so 1 followed by 100 zeroes in binary == 2^100
	Small = Big >> 99
	// Shift it right 99 places, so 1 followed by 1 zero in binary == 2^1 == 2
)

func needInt(x int64) int64 { 
	return x*10 + 1 
}

func needFloat(x float64) float64 {
	return x * 0.1
}



func main() {

	fmt.Println(needInt(Small))


	//fmt.Println(needInt(Big)) 
	// This will throw an error of overflow because Big is too large to be stored in an int, int is 32 bits in size - cannot use Big (untyped int constant 1267650600228229401496703205376) as int64 value in argument to needInt (overflows)


	fmt.Println(needFloat(Small))
	
	
	fmt.Println(needFloat(Big)) 
	// This will not throw an error because Big is a float64, which can store a much larger number than an int64

}
