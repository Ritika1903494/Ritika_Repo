package main

import "fmt"

type Address struct {
	Name string
	city string
	Pincode int
}

func main() {

	var a Address
	fmt.Println(a)
	a1 := Address{"ritika", "hoshiarpur", 146024}

	fmt.Println("Address1: ", a1)

	a2 := Address{Name: "Disha", city: "hoshiarpur",Pincode:146024}

	fmt.Printf("Name: %v,city: %v,Pincode: %v",a2.Name,a2.city,a2.Pincode);

	user1 := new(Address)
   user1.Name = "nargis"
   user1.city = "sultanpur"
   user1.Pincode = 25000
    fmt.Println(user1)
  
    
    user2 := new(Address)
   user2.Name = "Aman"
   user2.city = "kapurthala"
   user2.Pincode = 10012
    fmt.Println(user2)

	user3:= &Address{Name: "tarun", city: "hoshiarpur",Pincode:146024}
	fmt.Printf("Name: %v,city: %v,Pincode: %v",(*user3).Name,(*user3).city,(*user3).Pincode);
  
}
