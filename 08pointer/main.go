package main

import "fmt"
// pointer in go lang is same like c++
func main() {
	fmt.Println("pointer in golang")
	var ptr *int     // same like c and c++
	fmt.Println(ptr) // output <nil> because no addres is stored
	mynumber := 23
	var ptr2 = &mynumber     // we are using reference the value of mynumber
	fmt.Println(ptr2, *ptr2) //output will print address of variable and value of mynumber also
	*ptr2 = *ptr2 * 8
	fmt.Println(*ptr2) // output will be 23*8{ans khud nikal looo}
}
