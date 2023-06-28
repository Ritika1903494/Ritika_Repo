package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
)

type Employee struct {
	XMLName xml.Name `xml:"employee"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Email   string   `xml:"email"`
	Phone   string `xml:"phone"`
}

func main() {

	e1 := &Employee{Id: 101, Name: "Donald Duck", Email: "dd@gmail.com",Phone:"12345"}
	data, err := xml.MarshalIndent(e1, " ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile("emp.xml", data, 0666)
	if err != nil {
		log.Fatal(err)
	}

	data, err = ioutil.ReadFile("emp.xml")
	if err != nil {
		log.Fatal(err)
	}
	e1 = &Employee{}
	err = xml.Unmarshal([]byte(data), &e1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(e1.Id)
	fmt.Println(e1.Name)
	fmt.Println(e1.Email)
	fmt.Println(e1.Phone)
}