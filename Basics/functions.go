package main
import (
	"fmt"
)

func add(num1 int, num2 int) int {
	return num1 + num2
}

//shortened version (of above func) for sub
func sub(num1, num2 int) int {
	return num1 - num2
}

//swap function
func swap(a, b string) (string, string){
	return b, a
}


func main(){
	fmt.Println("Sum of 2 and 3 is", add(2, 3))
	fmt.Println("Sum of 5 and 7 is", add(5, 7))
	fmt.Println("Sum of 10 and 20 is", add(10, 20))

	//shortened version for sub
	fmt.Println("Difference of 10 and 5 is", sub(10, 5))
	fmt.Println("Difference of 20 and 10 is", sub(20, 10))
	fmt.Println("Difference of 100 and 50 is", sub(50, 150))

	//swap function
	a, b := swap("Tohed", "Anwar")
	fmt.Println(a, b)

}