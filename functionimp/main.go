package main

import (
	b "functionimp/basic"
	gross "functionimp/gross"

	"fmt"
)

func main() {
	b.Basic = 10000
	fmt.Println(gross.GrossSalary())
}