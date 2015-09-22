package db

import (
    "database/sql"
    _ "github.com/lib/pq"
    "errors"
)

func GetAllUsers() (users []User, err error) {
    db, err := sql.Open("postgres", "user=postgres password=postgres dbname=teste sslmode=disable")
    if err != nil {
        return nil, err
    }

    rows, err := db.Query("SELECT * FROM users")
    defer rows.Close()
    if err != nil {
        return nil, err
    }

    for rows.Next() {
        var name string
        var id, age int

        if err = rows.Scan(&id, &name, &age); err != nil {
            return nil, err
        }
        users = append(users, User{Id:id, Age:age, Name:name})
    }

    if err = rows.Err(); err != nil {
        return nil, err
    }

    return users, nil
}

func GetUserById(id int) (User, error) {
    db, err := sql.Open("postgres", "user=postgres password=postgres dbname=teste sslmode=disable")
    if err != nil {
        return User{}, err
    }

    var idR, age int
    var name string

    err = db.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&idR, &name, &age)
    switch {
    case err == sql.ErrNoRows:
        return User{}, errors.New("No user with that ID.")
    case err != nil:
        return User{}, err
    default:
        return User{Id: idR, Name: name, Age: age}, nil
    }
}

func DeleteUserById(id int) error {
    db, err := sql.Open("postgres", "user=postgres password=postgres dbname=teste sslmode=disable")
    if err != nil {
        return err
    }

    if _, err = db.Exec("DELETE FROM users WHERE id = $1", id); err != nil {
        return err
    }

    return nil
}

func SaveUser(name string, age int) error {
    db, err := sql.Open("postgres", "user=postgres password=postgres dbname=teste sslmode=disable")
    if err != nil {
        return err
    }

    _, err = db.Exec("INSERT INTO users(name, age) VALUES ($1, $2)", name, age)

    if err != nil {
        return err
    }

    return nil
}