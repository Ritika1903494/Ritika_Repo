package main

import "fmt"

func main() {
    
   defer fmt.Println("One")
   defer fmt.Println("Two")
   defer fmt.Println("Three")
   fmt.Println("Defer")
   division(4, 2)
   division(8, 0)
   division(2, 8)
}

func handlePanic() {
  /* detect if panic occurs or not*/
  a := recover()
  if a != nil {
    fmt.Println("RECOVER", a)
  }

}

func division(num1, num2 int) {

  /* execute the handlePanic even after panic occurs**/
  defer handlePanic()
  if num2 == 0 {
    panic("Cannot divide a number by zero")
  } else {
    result := num1 / num2
    fmt.Println("Result: ", result)
  }
}

