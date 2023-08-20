package main
import(
  //"fmt"
  "log"
  "net/http"
  "encoding/json"
)
type BasicData struct{
  Lang string `json:"Lang"`
  IntVal int32 `json:"numeric"`
  RandomText string `json:"rtext"`
}
func basicData(w http.ResponseWriter, r *http.Request){
  data := BasicData{Lang: "Go",IntVal: 32,RandomText: "Idk somethink"}
//  fmt.Fprintf(w,"Just hited the basicData endpoint")
  json.NewEncoder(w).Encode(data)

}

func handleRequests(){
  http.HandleFunc("/basic-data",basicData)
  log.Fatal(http.ListenAndServe(":8080",nil))
}

func main(){
  handleRequests()
}
