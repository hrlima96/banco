package utilities

import (
    "net/http"
    "github.com/hrlima96/banco/db"
    "encoding/json"
    "strconv"
)

func Users(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        respondUsers(w)
    case "POST":
        respondSaveUser(w, r)
    default:
        w.WriteHeader(http.StatusMethodNotAllowed)
    }
}

func respondUsers(w http.ResponseWriter) {
    users, err := db.GetAllUsers()
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    json, err := json.Marshal(users)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    jsonResponse(w, string(json))
}

func respondSaveUser(w http.ResponseWriter, r *http.Request) {
    name := r.FormValue("name")
    ageStr := r.FormValue("age")

    age, _ := strconv.Atoi(ageStr)
    
    if err := db.SaveUser(name, age); err != nil {
        w.WriteHeader(422)
        return
    }

    w.WriteHeader(http.StatusCreated)
}