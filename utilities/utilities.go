package utilities

import (
    "net/http"
    "fmt"
    "strings"
    "strconv"
    "log"
)

func jsonResponse(w http.ResponseWriter, json string) {
    log.Println("Rendering JSON")

    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintf(w, json)
}

func extractId(path string) (id int) {
    splitedPath := strings.Split(path, "/")
    idStr := splitedPath[len(splitedPath) - 1]
    id, _ = strconv.Atoi(idStr)

    return id
}

func getIP(r *http.Request) string {
    addr := r.RemoteAddr
    addrSplit := strings.Split(addr, ":")

    if addrSplit[0] == "[" {
        return "127.0.0.1"
    }

    return addrSplit[0]
}