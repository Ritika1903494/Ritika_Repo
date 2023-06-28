package main

import "fmt"
/* Public variable always start with capital*/
const Name="ritika";
func main() {
    fmt.Println("Constant variable")
	var fullName=Name+"thakur";
	fmt.Printf("the type of fullName is %T\n",fullName)
	fmt.Println("fullname:",fullName)
}