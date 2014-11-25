package controllers

import (
    "fmt"
    "net/http"
    "html/template"
)

func IndexRegister() {
    http.HandleFunc("/", indexAction)
    http.HandleFunc("/teest", testAction)
}

func indexAction(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("views/index/index.html")
    t.Execute(w, nil)
}

func testAction(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "TEST")
}