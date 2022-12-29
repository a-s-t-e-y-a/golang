package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var new_user string = "welcome to golang enter new string"
	fmt.Printf(new_user)

	reader := bufio.NewReader(os.Stdin)
	result, _ := reader.ReadString('\n')
	fmt.Printf(result)

	/*
		WE CANT USE ANY INPUT FROM USER AS ANOTHER TYPE BECUASE IT
		IS ALWAYS IN STRING FORMAT
		SO WE NEED TO CONVERISON TOPIC
	*/
	conversion, err := strconv.ParseFloat(strings.TrimSpace(result), 64) /// we user error comma syntax here also
	/*
		WE ACTUALLY USE AN ANOTHER BUILT IN LIBRARY
		CALLED AS strconv WHICH USES ParseFloat method
		which convert string into a float
	*/

	/*

		Lets implemment an error handler
		using if and else
	*/
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("there is no error")
		fmt.Println(conversion + 1)
		/*
			now you are thinking that conversion is now done
			you gonna enter a sting and it gonna convert to a string
			but wait that is no the exactly case going on here

			when you give a buffer to the terminal what gonna happen is
			"<your given string > /n" you will get /n with the string
			how you gonna convert a /n to float
			the ans we dont have to we just have trim the /n so lets do this

			BY USING A ANOTHER BUILT IN LIBRARY KNOW AS "strings" AND HAVE
			A METHOD CALLED AS "TrimSpace" SO IT GOES LIKE
			strings.TrimSpace(<convert variable>)
		*/
	}

}
