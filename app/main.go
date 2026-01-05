package main

import (
    "fmt"
    "net/http"

    "github.com/KushalMohith18/go-cache-heavy-demo/app/pkg1"
    "github.com/KushalMohith18/go-cache-heavy-demo/app/pkg2"
    "github.com/KushalMohith18/go-cache-heavy-demo/app/pkg3"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, pkg1.Work()+pkg2.Work()+pkg3.Work())
    })
    http.ListenAndServe(":8080", nil)
}
