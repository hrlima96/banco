package utilities

import (
    "github.com/hrlima96/banco/db"
    "net/http"
    "encoding/json"
)

func User(w http.ResponseWriter, r *http.Request) {
    id := extractId(r.URL.Path)

    switch r.Method {
    case "GET":
        respondUser(w, id)
    case "DELETE":
        respondDeleteUser(w, id)
    default:
        w.WriteHeader(http.StatusMethodNotAllowed)
    }
}

func respondUser(w http.ResponseWriter, id int) {
    user, err := db.GetUserById(id)
    if err != nil {
        w.WriteHeader(http.StatusNotFound)
        return
    }

    jsonR, err := json.Marshal(user)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    jsonResponse(w, string(jsonR))
}

func respondDeleteUser(w http.ResponseWriter, id int) {
    err := db.DeleteUserById(id)
    if err != nil {
        w.WriteHeader(http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}