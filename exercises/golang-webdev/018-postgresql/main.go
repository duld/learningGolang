/*
connect to a postgresql
- start postgres in a terminal
- connect to the postgres using a valid uri
-
*/

package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123dubDub"
	dbname   = "bookshelf"
)

func main() {
	// Postgres connection string.
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// ping our DB to establish a connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("connected to db!!!")
	}

	rows, err := db.Query("SELECT * FROM public.inventory;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}

		fmt.Println(id, name)
	}

	// get any error encountered during itteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}
}
