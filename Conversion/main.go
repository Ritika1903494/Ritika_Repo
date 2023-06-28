package main

import "fmt"

func main() {
    fmt.Println("Variables")
	var num1 int=2;
	var num2 int=5;
	var sum int=num1+num2;
	fmt.Printf("Sum is of type  %T\n",sum)
	fmt.Println("Sum is =",sum)

	var str1 string="ritika";
	var str2 string="thakur";
	var str3 string =str1+str2;
	fmt.Printf("Name is of type  %T\n",str3)
	fmt.Println("Name is =",str3)

	var avg float32= float32(sum)/2;
	fmt.Printf("Average is of type  %T\n",avg)
	fmt.Println("Average  is =",avg)

	var isRadius bool=false;
	var radius int;
	var area float32;
	if isRadius {
		fmt.Println(" isRadius =",isRadius)
		area=2*3.14*float32(radius)*float32(radius);
		fmt.Println("Area  is =",area)
	}else{
		fmt.Println(" isRadius =",isRadius)
		radius=2;
		area=2*3.14*float32(radius)*float32(radius);
		fmt.Println("Area  is =",area)
	}
}