package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)
type BasicData struct{
  Lang string `json:"Lang"`
  IntVal int32 `json:"numeric"`
  RandomText string `json:"rtext"`
}
type ClientData struct{
  Id uint32 `json:"id"`
  Desc string `json:"desc"`
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

func basicData(w http.ResponseWriter, r *http.Request){
  data := BasicData{Lang: "Go",IntVal: 32,RandomText: "Idk somethink"}
//  fmt.Fprintf(w,"Just hited the basicData endpoint")
  json.NewEncoder(w).Encode(data)

}

func handleRequests(){
  http.HandleFunc("/basic-data",basicData)
  http.HandleFunc("/sdata", jsonDataCall)
  log.Fatal(http.ListenAndServe(":8080",nil))
}

func main(){
  handleRequests()
}
