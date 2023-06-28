package father

import (
	"fmt"
)
var Data=10; 

type Father struct {
	Name string
}

func (f Father) Data(name string) (string,int){
	f.Name = "Father : " + name
	return f.Name,Data;
}

func Increment(){
	fmt.Println("Data is incremented:",Data)
}