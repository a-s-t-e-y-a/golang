package main

import "fmt"

func main() {
	fmt.Println("leaning loop in go lang")

	day := []string{"sunday", "monday", "tuesday", "wednesday"}
	/*
		here we are creating a slice
	*/
	fmt.Print(day)

	/*
		go has only for but with varity of variation
	*/

	for d := 0; d < 4; d++ {
		fmt.Println("\n", day[d])
	}
	/*
		we can also use this loop method for the looping through slices

		but this is no the case
	*/
	for i := range day {
		fmt.Println(day[i])
	}
	for index, day := range day {
		fmt.Print(index, day)
	}
}
