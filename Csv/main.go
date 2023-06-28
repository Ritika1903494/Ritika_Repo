package main

import (
 "encoding/csv"
 "fmt"
 "os"
)

type user struct {
 Id           string
 Name         string
 Lastname     string
 Username     string
 Organization string
}

func main() {

 // some sample data
 usersToExport := []user{
  {
   Id:           "097812-abc",
   Name:         "Jack",
   Lastname:     "Daniels",
   Username:     "jd007",
   Organization: "Google",
  },
  {
   Id:           "097812-abc",
   Name:         "Jack",
   Lastname:     "Daniels",
   Username:     "jd007",
   Organization: "Google",
  },
 }

 // create the file to write the values
 file, err := os.Create("./data.csv")
 defer file.Close()
 if err != nil {
  panic(err)
 }

 // create csv writer object
 w := csv.NewWriter(file)
 defer w.Flush()

 csvHeader := []string{"id", "name", "lastname", "username", "organization"}
 var csvData [][]string
 csvData = append(csvData, csvHeader)

 for _, user := range usersToExport {
  csvData = append(csvData, []string{user.Id, user.Name, user.Lastname, user.Username, user.Organization})
 }

 w.WriteAll(csvData)
 Read_csv();

}

func Read_csv(){
var users []user
 file, err := os.Open("./data.csv")
 if err != nil {
  fmt.Println(err)
 }

 defer file.Close()

 reader := csv.NewReader(file)
 csvRecords, err := reader.ReadAll()

 if err != nil {
  panic(err)
 }

 for i := 1; i < len(csvRecords); i++ {
  users = append(users, user{
   Id:           csvRecords[i][0],
   Name:         csvRecords[i][1],
   Lastname:     csvRecords[i][2],
   Username:     csvRecords[i][3],
   Organization: csvRecords[i][4],
  })
 }

 fmt.Println(users)
}