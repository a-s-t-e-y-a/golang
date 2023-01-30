package main

import "fmt"

func main() {
	fmt.Println("structs in golang")
	// no inheritance in golang
	// no super or parent concept in go
	krishna := User{"krishna", "krishna.go.dev", "true", 16, 12}
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
	krishna.GetStatus()        // calling the method from here
	krishna.NewMail()          // calling the method from here
	fmt.Println(krishna.Email) // here we printing email agian
	// but it is still gonna return krishna.dev.go instead of new email that we give test@dev.com and now why this happen because it is passing as copy not original value
}

/*
how to define struct in go lang
*/
type User struct {
	Name   string
	Email  string
	Status string
	Age    int
	oneAge int
}

/*
PUT FIRST LETTER CAPITAL IN STRUCTURE{so they are exportable}
capital letter things are exportable and methos start with small letter are not exportable
*/

/*
we are creating a fucntion method here for the structs
so what we do is
we place bracket before the givin name to the method
{Getstatus} and remember one thing the first letter of the  name that you give keep the first letter capital
and pass u of the type user {which is the struct}
*/
func (u User) GetStatus() {
	fmt.Println("is user active", u.Status)
}
func (u User) NewMail() {
	u.Email = "test@dev.com"
	fmt.Println("email of the user", u.Email)
}
