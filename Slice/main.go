package main

import "fmt"

func main() {
    fmt.Println("Slice")
	array := [5]int{1, 2, 3, 4, 5}
	/* start number is included and last number is exclusive */
    slice := array[1:4]
 
    fmt.Println("Array: ", array)
    fmt.Println("Slice: ", slice)

	slice2:= []int{1, 2, 3}
    slice2 = append(slice2, 4, 5, 6)
 
    fmt.Println("Slice: ", slice2)
	/* access*/
	fmt.Println("slice",slice2[1])
	slice2=append(slice2[:4],slice2[4+1:]...)
		/* after removing the one element from slice2*/
	fmt.Println("Slice: ", slice2)
	slice3:=make([]string,5);

	for i:=0;i<len(slice3);i++{
		fmt.Println("Enter your name:")
		var name string
		fmt.Scanln(&name)
		slice3[i]=name
	}

	for index,value:=range slice3{
		fmt.Printf("index %v ,name is %v\n",index,value)
	
	}   

}