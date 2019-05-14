package main
import(
    "io"
    "net/http"
)
func main(){
    main_handler := func(w http.ResponseWriter, req *http.Request){
        io.WriteString(w, "hello,world")
    }
    http.HandleFunc("/hello", main_handler)
    http.ListenAndServe(":8080", nil)
}
