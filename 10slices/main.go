package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("slices in go lang")
	var fruitlist = []string{"apple", "banana", "gauva"}
	/*
		we can declare slices this where we intiliazing also simultaneously
		slices are not hardcore bound means they get expanded with the as the space allocation gets changed

		we will see it later in more chapter
	*/
	fmt.Printf("type of fruit list is %T\n", fruitlist)

	/*
		we will use append method to appned new elements to the slices in go lang
	*/
	fruitlist = append(fruitlist, "new fruit 1", "new fruit 2")
	fmt.Println(fruitlist)
	fruitlist = append(fruitlist[1:3])
	/*
		this will same work like javascript slices method
		and everything like javascript slices
	*/
	fmt.Println(fruitlist)

	/*

		WE ARE GOING TO USE MAKE KEYWORD TO DEFINE SLICES
		{make keyword that we use in the memory management}

	*/
	highscore := make([]int, 4)
	/*
		here we need to just understand the syntax
		make(type , size)
		if you exceed the size it will get crashed
	*/
	highscore[0] = 234
	highscore[1] = 945
	highscore[2] = 465
	highscore[3] = 867

	/*
		so here is the question can we use append method with make because we know if we leap the size of the make we can make our program get crashed

		BUT NO THAT'S NOT THE CASE WITH MAKE()
		WITH MAKE YOU CAN USE APPEND METHOD AND CAN INCREASE THE SIZE OR MAKE ALTERNATIVE CHANGES TO THE SLICES WITH MAKE METHOD
	*/
	highscore = append(highscore, 12, 13, 45)
	fmt.Println(highscore)
	sort.Ints(highscore) // sort the int in increasing order
	fmt.Println(highscore)
	fmt.Println(sort.IntsAreSorted(highscore))

	/*
		HOW TO REMOVE VALUE FROM SLICE BASED ON INDEX
	*/
	var courses = []string{"java", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	/*
		here we are trying to remove swift so we use append to remove
		yeah we can remove intems also by using append method
		and why we use (...) spread operator like javascript
		we are gonna look into next videos
	*/
	fmt.Println(courses)

}
