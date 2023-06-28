package main

import (
    "fmt"
    "strings"
)
  
const Num1=10;
const Num2=3;
const Num3=30;
const Num4=1;
const Num5=20;
const Num=5;
const Name="ritika"
func main() {
    fmt.Println("if else condition")
    if Num%2==0{
		fmt.Println( Num,"is even")
	}else{
		fmt.Println(Num,"is odd")
	}

	for i:=0;i<len(Name);i++{
         var ch=string(string(Name)[i]);
		 fmt.Println(ch)
		fmt.Println("count",strings.Count(Name,ch))	
	}

	if Num1>Num2 && Num1>Num3 && Num1>Num4 && Num1>Num5{
		fmt.Println( Num1,"is Greater Number")
	}else if Num2>Num3 && Num2>Num4 && Num2>Num5{
		fmt.Println( Num2,"is Greater Number")
	}else if Num3>Num4 && Num3>Num5{
		fmt.Println( Num3,"is Greater Number")
	}else if Num4>Num5{
		fmt.Println( Num4,"is Greater Number")
	}else{
		fmt.Println( Num5,"is Greater Number")
	}
}