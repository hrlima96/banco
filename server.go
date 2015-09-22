package main

import (
    "fmt"
    "log"
    "net/http"
    "flag"
    "github.com/hrlima96/banco/utilities"
)

var port *int

func init() {
    port = flag.Int("p", 8888, "port")
    flag.Parse()
}

func main() {
    http.HandleFunc("/users", utilities.Users)
    http.HandleFunc("/users/", utilities.User)

    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}