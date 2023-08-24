package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
  "database/sql"
  _"github.com/go-sql-driver/mysql"
)
var db *sql.DB 
type BasicData struct{
  Lang string `json:"Lang"`
  IntVal int32 `json:"numeric"`
  RandomText string `json:"rtext"`
}
type ClientData struct{
  Id uint32 `json:"id"`
  Desc string `json:"desc"`
}
type DBData struct{
  Id uint32 `json:"id"`
  Name string `json:"name"`
  LastName string `json:"last_name"`
}
func jsonDataCall(w http.ResponseWriter, r *http.Request){
  var cd ClientData
  err := json.NewDecoder(r.Body).Decode(&cd)
  if err != nil{
    http.Error(w , err.Error(), http.StatusBadRequest)
    return
  }
  fmt.Fprint(w , "Recieved: ",cd)
  if cd.Id == 42{
    fmt.Fprint(w ,"\nYou found the he meaning of life, the universe, and everything")
  }
}
func dataFromDb(w http.ResponseWriter, r *http.Request){


  if r.Method == http.MethodGet{
  res, err := db.Query("SELECT * FROM test_table")
  defer res.Close()
  if err!=nil{
    fmt.Fprint(w,"Error while reading DB:\n",err)
    return
  }
  fmt.Fprint(w,"{\n")
for res.Next(){
    var dat DBData
    err:=res.Scan(&dat.Id,&dat.Name,&dat.LastName)
    if err!=nil{
      return
    }
  json.NewEncoder(w).Encode(dat)
  }
  fmt.Fprint(w , "}")
  } else if r.Method == http.MethodPost{//TODO: test if works
    var data DBData
    err := json.NewDecoder(r.Body).Decode(data)
    if err != nil{
      fmt.Fprint(w,"ERROR: Unable to decode data")
      return 
    }
    v, err:= db.Query("INSERT INTO test_table VALUES(?, ?, ?)", data.Id,data.Name,data.LastName)
    if err!=nil{
      fmt.Fprint(w, "DB write failed")
    return
    }
    fmt.Fprint(w,"Data saved sucesfully")
  }
  }

func basicData(w http.ResponseWriter, r *http.Request){
  data := BasicData{Lang: "Go",IntVal: 32,RandomText: "Idk somethink"}
//  fmt.Fprintf(w,"Just hited the basicData endpoint")
  json.NewEncoder(w).Encode(data)

}

func handleRequests(){
  http.HandleFunc("/basic-data",basicData)
  http.HandleFunc("/sdata", jsonDataCall)
  http.HandleFunc("/rdb", dataFromDb)
  log.Fatal(http.ListenAndServe(":8080",nil))
}

func main(){
  var err error
  db, err =sql.Open("mysql", "root@tcp(127.0.0.1:3306)/go_db")
  if err != nil{
    log.Fatal(err)
  }
  handleRequests()
defer db.Close()
}
