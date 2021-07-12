package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:12345@/taxi_service")
	log.Fatal(err)

	log.Println("DB Connected", db)

	// defer db.Close()
}
