package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123456"
	dbname   = "postgres"
)

type Student struct {
	ID   int
	Name string
}

func main() {
	user := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", user)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	sqlStatement := `SELECT * FROM student WHERE id=$1;`
	var student Student
	row := db.QueryRow(sqlStatement, 1)
	err = row.Scan(&student.ID, &student.Name)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return
	case nil:
		fmt.Println(student)
	default:
		panic(err)
	}
}
