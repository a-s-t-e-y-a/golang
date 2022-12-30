/*
array is the least used in the golang in comparrision to other
language
*/
package main

import "fmt"

func main() {
	fmt.Println("array in golang")
	var array [4]string
	array[0] = "apple"
	array[1] = "mango"
	array[3] = "strawberry"
	fmt.Println(array)
	/*
	 output will be
	 [apple mango  strawberry]
	 notice the space between mango and strawberry
	 why? becuase no value given to the index 2
	 so golnag  try to give you the hint
	*/
	fmt.Println(len(array))

	/*
		after all we have only two value in the array but
		it will show len as 4
	*/
	var veg = [3]string{"potato", "benas", "mushroom"}
	fmt.Println(veg)
	/*
		you can define array as this also
	*/

}
