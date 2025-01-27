package main

import "fmt"

func main(){
	v := 42
	fmt.Printf("v is of type %T", v)

	//		you can not assign different type into v

	// v := 45.6
	// fmt.Printf("v is of type %T", v)
	// v := 50        // ERROR:  (you can not redeclare v)
	// fmt.Printf("\nv is of type %T", v)

	//above code will throw an error


	//		if you want to assign different values (of same type), you can assign
	v = 45
	fmt.Printf("\nv is of type %T", v)

	var i int
	j := i // j is an int (inferred from i)

	//more type inference
	k := 42           // int
	l := 3.142        // float64
	m := 0.867 + 0.5i // complex128
}