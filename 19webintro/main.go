package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("llco webiste hit")
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is of type: %T\n", res)
	defer res.Body.Close()                     //it is your duty to close the response of the body
	databytes, err := ioutil.ReadAll(res.Body) //PART OF FILE HABDLING IN THE GO
	if err != nil {
		panic(err)
	}

	content := string(databytes)
	fmt.Print(content)
}
