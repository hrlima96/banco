// Package db provides access the database teste
package db

import (
    "database/sql"
    _ "github.com/lib/pq"
    "errors"
    "fmt"
    "log"
)

type User struct {
    Id int `json:"id"`
    Name string `json:"name"`
    Age int `json:"age"`
}

func GetAllUsers() (users []User, err error) {
    db, err := sql.Open("postgres", "user=postgres password=postgres dbname=teste sslmode=disable")
    if err != nil {
        log.Println(err)
        return nil, err
    }

    query := "SELECT * FROM users ORDER BY id"
    log.Println(query)

    rows, err := db.Query(query)
    defer rows.Close()
    if err != nil {
        log.Println(err)
        return nil, err
    }

    for rows.Next() {
        var name string
        var id, age int

        if err = rows.Scan(&id, &name, &age); err != nil {
            log.Println(err)
            return nil, err
        }
        users = append(users, User{Id:id, Age:age, Name:name})
    }

    if err = rows.Err(); err != nil {
        log.Println(err)
        return nil, err
    }

    return users, nil
}

func GetUserById(id int) (User, error) {
    db, err := sql.Open("postgres", "user=postgres password=postgres dbname=teste sslmode=disable")
    if err != nil {
        log.Println(err)
        return User{}, err
    }

    var idR, age int
    var name string

    query := fmt.Sprintf("SELECT * FROM users WHERE id = %d", id)
    log.Println(query)

    err = db.QueryRow(query).Scan(&idR, &name, &age)
    switch {
    case err == sql.ErrNoRows:
        err = errors.New("No user with that ID.")
        log.Println(err)

        return User{}, err
    case err != nil:
        log.Println(err)
        return User{}, err
    default:
        return User{Id: idR, Name: name, Age: age}, nil
    }
}

func DeleteUserById(id int) error {
    db, err := sql.Open("postgres", "user=postgres password=postgres dbname=teste sslmode=disable")
    if err != nil {
        log.Println(err)
        return err
    }

    query := fmt.Sprintf("DELETE FROM users WHERE id = %d", id)
    log.Println(query)

    if _, err = db.Exec(query); err != nil {
        log.Println(err)
        return err
    }

    return nil
}

func SaveUser(user User) error {
    db, err := sql.Open("postgres", "user=postgres password=postgres dbname=teste sslmode=disable")
    if err != nil {
        log.Println(err)
        return err
    }

    query := fmt.Sprintf("INSERT INTO users(name, age) VALUES('%s', %d)", user.Name, user.Age)
    log.Println(query)

    _, err = db.Exec(query)
    if err != nil {
        log.Println(err)
        return err
    }

    return nil
}

func UpdateUserById(user User) (err error){
    db, err := sql.Open("postgres", "user=postgres password=postgres dbname=teste sslmode=disable")
    if err != nil {
        log.Println(err)
        return err
    }

    query := fmt.Sprintf("UPDATE users SET name = '%s', age = %d WHERE id = %d", user.Name, user.Age, user.Id)

    _, err = db.Exec(query)
    if err != nil {
        log.Println(err)
        return err
    }

    return nil
}