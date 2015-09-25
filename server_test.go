package main

import (
    "testing"
    "net/http"
    "github.com/hrlima96/banco/db"
    "encoding/json"
    "bytes"
)

func TestNonExistingPath(t *testing.T) {
    nonExistentPaths := []string{"http://localhost:8888/", "http://localhost:8888/users/5/"}

    for _, path := range nonExistentPaths {
        response, _ := http.Get(path)
        response.Body.Close()

        st := response.StatusCode
        if st != 404 {
            t.Errorf("Status errado, deveria ser 404")
        }
    }
}

func TestGetAllUsers(t *testing.T) { /* GET /users */
    users_db, _ := db.GetAllUsers()
    user_qtd_db := len(users_db)

    response, _ := http.Get("http://localhost:8888/users")
    defer response.Body.Close()

    var users_req []db.User

    _ = json.NewDecoder(response.Body).Decode(&users_req)
    user_qtd_json := len(users_req)

    if user_qtd_db != user_qtd_json {
        t.Errorf("Numero de users diferente")
    }
}

func TestSaveUser(t *testing.T) { /* POST /users */
    users_db, _ := db.GetAllUsers()
    user_qtd_db := len(users_db)

    json_user, _ := json.Marshal(db.User{Id: 13, Name: "Disney", Age: 20})

    req, _ := http.NewRequest("POST", "http://localhost:8888/users", bytes.NewBuffer(json_user))

    client := &http.Client{}
    _, _ = client.Do(req)

    users_db2, _ := db.GetAllUsers()
    user_qtd_db2 := len(users_db2)

    if user_qtd_db == user_qtd_db2 {
        t.Errorf("Numero de users igual, deveria ter aumentado!")
    }
}