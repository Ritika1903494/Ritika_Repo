package main

import "fmt"

func main() {
    fmt.Println("Function")
	calculator(2,4,"*")
	area := func(length, breadth int) int {
		return length * breadth
	  } 
	
	  fmt.Println("The area of rectangle is", area(3,4))

	  var area1, area2 = myfunc(2, 4)
	

fmt.Printf("Area of the rectangle is: %d", area1 )
fmt.Printf("\nArea of the square is: %d", area2)
	
}

func calculator(num1 int,num2 int,operator string){
	if operator=="+"{
		fmt.Println("addition is :",num1+num2)
	}else if  operator=="-"{
		fmt.Println("subtraction  is :",num1-num2)
	}else if  operator=="*"{
		fmt.Println("multiplication  is :",num1*num2)
	}else{
		fmt.Println("division  is :",int32(num1/num2))
	}

}



func myfunc(p, q int)( rectangle int, square int ){
	rectangle = p*q
	square = p*p
	return
}




