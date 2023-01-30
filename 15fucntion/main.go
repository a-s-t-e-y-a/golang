package main

import "fmt"

func main() {
	fmt.Println("learnig fucntion in go lang")
	/*
		main is also an function and it acts like a entry point in the programm
	*/
	greeter() // calling a fucntion

	/*
		you can not write function inside a function
	*/
	/*
		anonymus function and iif fucntion also exist
	*/
	q := greeterTwo(2, 4)
	fmt.Println(q)
	greeterThree(1, 2, 3, 4, 5, 6, 7, 8, 9)
}

/*
decalring a fucntion
*/
func greeter() {
	fmt.Println("namsety form go lang")
}

/*
FUNCTION THAT ACCEPT ARGUMENT AND RETURN A VALUE
THE INT THAT WE WRITE BEFORE CURLY BRACKET IS RETURN TYPE
YIU HAVE TO DEFINE VARIBALE TYPE IN ARGUMENT
*/
func greeterTwo(a int, b int) int {
	fmt.Println("namsety form go lang", a+b)
	return a + b
}

/*
HERE we are passing many int value in the function so that we can spread them up
*/

func greeterThree(a ...int) (int, string) {
	sum := 0
	/*
		here we are looping throiugh the value that we pass form the fucntion
	*/

	for i := range a {

		sum = sum + a[i]
		fmt.Println("sum", sum)
	}
	fmt.Println("namsety form go lang", sum)
	return sum, "new value" // you can return two type from a single
	// function just need to ad dnew datatype at ending
	// like int and string

}
