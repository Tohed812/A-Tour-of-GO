package main

import (
	"fmt"
)

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum + 4 - 9
	return 
}

func main(){
	fmt.Println(split(17))
}

