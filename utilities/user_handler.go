package utilities

import (
    "github.com/hrlima96/banco/db"
    "net/http"
    "encoding/json"
    "log"
)

func User(w http.ResponseWriter, r *http.Request) {
    id := extractId(r.URL.Path)
    clientIP := getIP(r)

    switch r.Method {
    case "GET":
        log.Printf("Started GET to /users/%d for %s", id, clientIP)
        respondUser(w, id)
    case "DELETE":
        log.Printf("Started DELETE to /users/%d for %s", id, clientIP)
        respondDeleteUser(w, id)
    case "PUT":
        log.Printf("Started PUT to /users/%d for %s", id, clientIP)
        respondUpdateUser(w, r, id)
    default:
        w.WriteHeader(http.StatusMethodNotAllowed)
    }
}

func respondUser(w http.ResponseWriter, id int) {
    user, err := db.GetUserById(id)
    if err != nil {
        log.Println(err)
        w.WriteHeader(http.StatusNotFound)
        return
    }

    jsonR, err := json.Marshal(user)
    if err != nil {
        log.Printf("%d %v", http.StatusInternalServerError, err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    jsonResponse(w, string(jsonR))
    log.Printf("200 OK")
}

func respondDeleteUser(w http.ResponseWriter, id int) {
    err := db.DeleteUserById(id)
    if err != nil {
        log.Printf("%d %v", http.StatusNotFound, err)
        w.WriteHeader(http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}

func respondUpdateUser(w http.ResponseWriter, r *http.Request, id int) {
    var user db.User

    decoder := json.NewDecoder(r.Body)
    err := decoder.Decode(&user)
    if err != nil {
        log.Printf("%d %v", 422, err)
        w.WriteHeader(422)
        return
    }

    if err := db.UpdateUserById(user); err != nil {
        log.Printf("%d %v", http.StatusNotFound, err)
        w.WriteHeader(http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusOK)
}