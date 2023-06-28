package main

import "fmt"

func main() {
    fmt.Println("Pointer")
	var ptr *int;
	mynumber:=23;
	ptr=&mynumber;
	fmt.Println("Pointer memory location:",&ptr)
	fmt.Println("Pointer store memory address of variable:",ptr)
	fmt.Println("to access value using pointer:",*ptr)
}