
package main

import "fmt"

type tank interface {

	area() float64
	circumference() float64
}

type myvalue struct {
	radius float64
}


func (m myvalue) area() float64 {

	return 3.14*m.radius * m.radius;
}

func (m myvalue) circumference() float64 {

	return 2*3.14* m.radius;
}


func main() {
	
	var t tank
	t = myvalue{10}
	fmt.Println("Area of tank :", t.area())
	fmt.Println("Volume of tank:", t.circumference())
}
