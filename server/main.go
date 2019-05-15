package main
import(
    "io"
    "io/ioutil"
    "net/http"
)
func main(){
    main_handler := func(w http.ResponseWriter, req *http.Request){
        if content, err := ioutil.ReadFile("main.html"); err != nil{
            io.WriteString(w, "read main.html error")
        }else{
            w.Write(content)
        }
    }
    http.Handle("/css/", http.FileServer(http.Dir("./")))
    http.Handle("/js/", http.FileServer(http.Dir("./")))
    http.Handle("/html/", http.FileServer(http.Dir("./")))
    http.Handle("/video/", http.FileServer(http.Dir("./")))
    http.HandleFunc("/", main_handler)
    http.ListenAndServe(":8080", nil)
}
