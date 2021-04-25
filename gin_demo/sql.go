package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := sql.Open("mysql",
		"root:123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal(err)
		fmt.Println("fail.........")
	}
	defer db.Close()

}
