package main

import (
    "fmt"
    "flag"
    "github.com/stdclass/go-mvc/controllers"
    "net/http"
)

var port = flag.Int("port", 8080, "http service port")

func main() {
    
    flag.Parse()
    
    controllers.IndexRegister()
    //controllers.NewsRegister()
    
    http.ListenAndServe(":8080", nil)
    
    
    fmt.Println("")
}