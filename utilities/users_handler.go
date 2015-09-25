package utilities

import (
    "net/http"
    "github.com/hrlima96/banco/db"
    "encoding/json"
    "log"
)

func Users(w http.ResponseWriter, r *http.Request) {
    clientIP := getIP(r)

    switch r.Method {
    case "GET":
        log.Printf("Started GET to /users for %s", clientIP)
        respondUsers(w)
    case "POST":
        log.Printf("Started POST to /users for %s", clientIP)
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
    var user db.User

    decoder := json.NewDecoder(r.Body)
    err := decoder.Decode(&user)
    if err != nil {
        log.Println("Unprocessable entity 422")
        w.WriteHeader(422)
        return
    }

    log.Printf("params: %+v", user)

    if err := db.SaveUser(user); err != nil {
        log.Println("Unprocessable entity 422")
        w.WriteHeader(422)
        return
    }

    w.WriteHeader(http.StatusCreated)
}