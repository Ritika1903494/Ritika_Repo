package main

import "fmt"
const Temp=30;
func main() {
    fmt.Println("Switch")
	var mark int=85;
	fmt.Println("convert temperature in : ")
    var second string
    fmt.Scanln(&second)

    switch {
	case mark>=90:
		fmt.Println("Grade is A+")
	case mark>=80:
		fmt.Println("Grade is A")
    case mark>=70:
		fmt.Println("Grade is B+")
	case mark>=60:
		fmt.Println("Grade is C")
	default:
		fmt.Println("Fail")
	}

	var convert=func(convert string){
		switch{
		case convert=="celsius":
			var cel float32=(Temp-32)*5/9;
			fmt.Println("temperature in celsius:",cel)
	    
		case convert=="fahrenheit":
			var fah float32=(Temp*1.8)+32;
			fmt.Println("temperature in fahrenheit:",fah)
	
		default:
			fmt.Println("Fail to convert ")
		}
	}

	convert(second)

}