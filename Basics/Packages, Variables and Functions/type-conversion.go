package main 

import (
	"fmt"
	"math"
)

var (
	i, j int = 42, 2701
	f float64 = float64(i) //converted to float64
	u uint = uint(f) //converted to uint
) 

func main(){
	fmt.Println(i, f, u)
	
	var ans float64 = math.Sqrt(float64(i*i+j*j))
	fmt.Println(ans)
}