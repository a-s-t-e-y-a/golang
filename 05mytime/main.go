package main

import (
	"fmt"
	"time"
)

/*
WE ARE GOING TO USE
TIME PACKAGE
*/
func main() {
	fmt.Println("lets start learning time")
	presenttime := time.Now() // we are using a built in libaray known as time
	fmt.Println(presenttime)
	/*
		this will give output as{output is when i am using it can change with time}
		2022-12-30 04:23:34.087325259 +0530 IST m=+0.000051117
	*/

	/*
		but i dont want like this
		we need to format it like we did in javascript
	*/

	fmt.Println(presenttime.Format("01-02-2006"))
	/*
		WE ARE NOW GONNA USE ANOTHER LIBRARY
		"Format" lets it will format the time

		so it will give output like
		12-30-2022

	*/
	fmt.Println(presenttime.Format("01-02-2006 15:04:05 Monday")) // you can add some more parameter to Format like day,time
	/*
		12-30-2022 04:32:55 Friday
		{output gets replaced by current date}

	*/

	/*
		NOTE:-
		BY GIVING A FORMAT IN Format which is ""
		you  have to write exact thing to get the output replaced by that current time

		for example agar humne Monday likha hai toh output mein Monday replace ho jayega uss din ki date se
		and same thing with the time also ki agar humne current time to change karna hai toh humne
		15:04:05 hee likhna padega and if we write something else
		it not gonna work
		woh as it is ka as it is utha ke print kardega

	*/
	createddate := time.Date(2020, time.August, 10, 23, 23, 0, 0, time.UTC)
	/*
		WE ARE CRETAING A CUSTOME DATE HERE
		IT WORKS LIKE
		time.Date(year,month{always using time.<month>},day,hour,minute,second,nano second , location)
	*/
	fmt.Println(createddate)

}
