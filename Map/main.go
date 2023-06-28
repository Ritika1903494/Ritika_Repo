package main

import (
    "fmt"
    "sort"
)
func main() {
    fmt.Println("Map")
	map1 := map[int]string{
     
		1: "Dog",
		2: "Cat",
		3: "Cow",
		4: "Bird",
		5: "Rabbit",
}

for id, pet := range map1 {
	fmt.Println(id, pet)
}

delete(map1, 3)
delete(map1, 5)

for id, pet := range map1 {
	fmt.Println(id, pet)
}
var employee = map[int]string{}

for i:=0;i<4;i++{
	fmt.Println("Enter employee name:")
	var name string
	fmt.Scanln(&name)
	employee[i]=name;
} 

for key, name := range employee {
	fmt.Println(key,name)
}

basket := map[string]int{"orange": 5, "apple": 7,
"mango": 3, "strawberry": 9}

  keys:= make([]string, 0, len(basket))
 
    for k := range basket{
        keys = append(keys, k)
    }
    sort.Strings(keys)
 
    for _, k := range keys {
        fmt.Println(k, basket[k])
    }

	merge(map1,employee)


}

func merge(map1 map[int]string,employee map[int]string){
  
	for key, value:= range map1 {
         employee[key]=value
    }
	for _, k := range employee {
        fmt.Println(k)
    }
}