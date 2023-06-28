package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Sign struct{
	UserName   string   `json:"userName"`
	Email      string   `json:"email"` 
	Password   string    `json:"password"`       
}


func main() {
    fmt.Println("Crud operation")
	// Get_operation()
	// Post_operation()
	// Patch_operation()
	// Put_operation()
	Delete_operation()
}

func Check_error(err error){
      panic(err)
}

func Get_operation(){
    client := &http.Client{}
    req, err := http.NewRequest("GET", "http://localhost:3000/SignUp", nil)
    if err != nil {
       Check_error(err)
    }
   res, err1 := client.Do(req)
    if err1 != nil {
		Check_error(err1)
    }
    defer res.Body.Close()
    body, err2 := ioutil.ReadAll(res.Body)
    if err2 != nil {
		Check_error(err2)
    }
    fmt.Println(string(body))
}

func Post_operation(){
	user:=Sign{
	   UserName: "Amanjit kaur",
       Email:"Aman@gmail.com",
       Password:"A@123man",
	  }
	  body, _ := json.Marshal(user)
	  fmt.Println(body)
	  resp,err := http.Post("http://localhost:3000/SignUp","application/json",bytes.NewBuffer(body))
	  fmt.Println(bytes.NewBuffer(body))
	  if err != nil {
		  panic(err)
	  }
	  defer resp.Body.Close()
	  if resp.StatusCode == http.StatusCreated {
		  fmt.Println(http.StatusCreated)
		  body, err := ioutil.ReadAll(resp.Body)
		  if err != nil {
			  panic(err)
		  }
		  jsonStr := string(body)
		  fmt.Println("Response: ", jsonStr)
  
	  } else {
		  fmt.Println("Get failed with error: ", resp.Status)
	  }
	 
}

func Patch_operation(){
	url := "http://localhost:3000/SignUp/2"
	payload := map[string]string{"userName": "Tarun Kumar"}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}


func Put_operation(){
	url := "http://localhost:3000/SignUp/2"
    data := []byte(`{"userName": "John", "email":"Johan@gmail.com"}`)
    req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(data))
    if err != nil {
       panic(err)
    }

    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
       panic(err)
    }
    defer resp.Body.Close()
}

func Delete_operation(){
	client := &http.Client{}
    url := "http://localhost:3000/SignUp/2"
    req, err := http.NewRequest("DELETE",url, nil)
    if err != nil {
        panic(err)
    }

    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
}