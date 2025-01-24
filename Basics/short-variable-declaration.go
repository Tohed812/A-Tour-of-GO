package main

import "fmt"

func main(){
	var a, b int = 1, 2 //var declaration with data type
	
	var c, d = true, "Tohed" //dynamic variable (without explicit data type declaration)
	
	ts, python, golang := true, false, "yes!" //short variable declaration (without var keyword)

	fmt.Println(a, b, c, d, ts, python, golang)
}