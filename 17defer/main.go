package main

import "fmt"

func main() {
	fmt.Println("hello")
	defer fmt.Println("world")
	/*
		after tun this it will give you the same result
		that we are expecting
	*/
	defer fmt.Println("world")
	fmt.Println("hello")
	/*
		but when we are interchnaging line no 6 and 7 what gonna
		happen ..
		as soon as the defer encounter in the line no 12
		the {for imagination only}
		the go cuts the line 12 and place just before the
		ending curly braces ...
		and so that line 12 execute at the very last

	*/
	/*
		if there were many deffred statemnet what gonna happen ..
		they will execute in the reverse order
	*/
	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three")
	defer fmt.Println("four")
	/*
		result will be four , three , two , one
	*/
	mydefer() //FUNCTION EXECUTION DONE AS SAME ..
	//but inside fucntion we use defer the output of fucntion will be
	// like 4321

}
func mydefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
