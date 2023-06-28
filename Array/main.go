package main

import "fmt"

func main() {
    fmt.Println("Array")
	var Array1=[5] string{"ritika","disha","tarun","aman","nargis"};
	var Array2=[5]int{};
	for i:=0;i<len(Array2);i++{
		fmt.Println("Enter your birth date:")
		var date int
		fmt.Scanln(&date)
		Array2[i]=date;
	}   
    
	for index,value:=range Array1{
		fmt.Printf("index %v ,name is %v\n",index,value)
	
	}   

	for _,value:=range Array2{
		fmt.Printf("date is %v\n",value)
	
	}   



}