package main

import (
	cir "interfaceimport/circle"
	cub "interfaceimport/cube"

	"fmt"
)

func main() {
	var cirinterface  cir.Circletype;
	var cubinterface  cub.Cubetype;

	cirinterface=cir.Myvalue{2};
	fmt.Println("Area of Circle :",cirinterface.Area())
	fmt.Println("Circumference of Circle:",cirinterface.Circumference())

	cubinterface=cub.Myvalue{2};
	fmt.Println("Area of Cube :",cubinterface.Area())
	fmt.Println("Permiter of Cube:",cubinterface.Permiter())
}