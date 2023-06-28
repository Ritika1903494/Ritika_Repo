package main

import (
	parent "family/father"
	child "family/daughter"
	"fmt"
)

func main() {
	f := new(parent.Father)
	fmt.Println(f.Data("Mr.Sohan Singh"))
	address:=&parent.Data;
	fmt.Println("address",address);
	fmt.Println("address",*address);
     *address++;
     parent.Increment();
	d:= new(child.Daughter)
	fmt.Println(d.Data("Ritika thakur"))
}