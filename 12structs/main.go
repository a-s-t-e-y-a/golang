package main

import "fmt"

func main() {
	fmt.Println("structs in golang")
	// no inheritance in golang
	// no super or parent concept in go
	krishna := User{"krishna", "krishna.go.dev", "true", 16}
	/*
		enter data as you defined in the struct
	*/
	fmt.Print(krishna)
	/*
		if you use +v in the printf you gonna have
		all the deatils of that struncture you defined
	*/
	fmt.Printf("krishna details are : %+v\n", krishna)
	fmt.Printf("krishna details are : %+v\n", krishna.Name)
	/*
		we can also access particular value from the srtucture
		like we are going to use only name form the field
	*/
}

/*
how to define struct in go lang
*/
type User struct {
	Name   string
	Email  string
	Status string
	Age    int
}

/*
	PUT FIRST LETTER CAPITAL IN STRUCTURE
*/
