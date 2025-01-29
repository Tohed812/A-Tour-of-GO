package main

import "fmt"

func main(){

	// Go has only one looping construct, the for loop
	sum := 1
	for i := 0; i < 10; i++{
		sum += i
	}
	fmt.Println(sum) 


	// The init and post statements are optional
	sum = 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)


	// Go's while loop - for is Go's while
	sum = 5
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	
	// Infinite loop
	// for {
	// 	fmt.Println("Infinite loop")
	// }
}