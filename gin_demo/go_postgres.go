package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123456"
	dbname   = "userinfo"
)

func connectDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
	return db
}

func insertUser(db *sql.DB) {
	stmt, err := db.Prepare("insert into sdbo_department(dep_id,dep_name) values($1,$2)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(1, "mgr")

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("insert into sdbo_department success!")
	}

}

func query(db *sql.DB) {
	var id, name string

	rows, err := db.Query(" select * from sdbo_department where dep_id=$1", "1")

	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &name)

		if err != nil {
			fmt.Println(err)
		}
	}

	err = rows.Err()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(id, name)
}

func main() {
	db := connectDB()
	insertUser(db)
	query(db)

}
