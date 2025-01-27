package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("My fav number", rand.Intn(10))             // rand.Intn(10) generates a random number between 0 and 10
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(6)) //from math/rand package
}
