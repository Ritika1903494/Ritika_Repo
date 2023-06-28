package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Users struct{
	Companyname string
	Employee []User
}
type User struct {
	Id       int
	Name     string
	Email    string
	Role     string
	Password string
}

func main() {
	users:= Users{
	Companyname:"Bebo technologies",
	Employee:[]User{
    { 
	Id : 101,
	Name:"ritika thakur",
	Email:"ritika@gmail.com",
	Role:"Web developer",
	Password:"rtika@123",
	},
	{
    Id : 102,
	Name:"Disha thakur",
	Email:"Disha@gmail.com",
	Role:"Java developer",
	Password:"Disha@123",
	 },
	{
	Id : 103,
	Name:"Nargisdeep kaur",
	Email:"Nargis@gmail.com",
	Role:"Java developer",
	Password:"Nargis@123",
	},
	},
	}
	content,err :=json.Marshal(users)
	if err != nil {
		Check_error(err)
	}
	err1 :=ioutil.WriteFile("./dbjson.json",content,0644)
	if err1 != nil {
		Check_error(err1)
	}
	Read_json()
}

func Check_error(err error){
	panic(err)
}

func Read_json(){
	file, _ := ioutil.ReadFile("./dbjson.json")
     data :=Users{}
	err2:=json.Unmarshal([]byte(file), &data)
	if err2 != nil {
		Check_error(err2)
	}else{
		fmt.Println(len(data.Employee))
		for i := 0;i<len(data.Employee);i++ {
			fmt.Println("Id:",data.Employee[i].Id)
			fmt.Println("Id:",data.Employee[i].Name)
			fmt.Println("Email:",data.Employee[i].Email)
			fmt.Println("Role:",data.Employee[i].Role)
			fmt.Println("Password:",data.Employee[i].Password)
		}
	}
	
}