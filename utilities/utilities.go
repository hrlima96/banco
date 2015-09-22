package utilities

import (
    "net/http"
    "fmt"
    "strings"
    "strconv"
)

func jsonResponse(w http.ResponseWriter, json string) {
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintf(w, json)
}

func extractId(path string) (id int) {
    splitedPath := strings.Split(path, "/")
    idStr := splitedPath[len(splitedPath) - 1]
    id, _ = strconv.Atoi(idStr)

    return id
}