package main
import(
  "fmt"
  "log"
  "net/http"
)

func basicData(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w,"Just hited the basicData endpoint")
}

func handleRequests(){
  http.HandleFunc("/basic-data",basicData)
  log.Fatal(http.ListenAndServe(":8080",nil))
}

func main(){
  handleRequests()
}
