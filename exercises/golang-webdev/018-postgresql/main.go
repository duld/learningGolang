/*
connect to a postgresql
- start postgres in a terminal
- connect to the postgres using a valid uri
-
*/

package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

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

	text := getUserInput()
	if len(text) < 3 {
		return
	}

	insertBook(db, text)
}

func insertBook(db *sql.DB, bk string) {
	// get value of last row
	id := getLastID(db)

	sqlStatement := `
	INSERT INTO public.inventory (id, name) VALUES ($1, $2);`

	_, err := db.Exec(sqlStatement, id+1, bk)
	if err != nil {
		panic(err)
	}
}

func getLastID(db *sql.DB) int {

	rows, err := db.Query("SELECT * FROM public.inventory ORDER BY id DESC LIMIT 1;")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var id int
	var name string

	for rows.Next() {

		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("The last id is: ", +id)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return id
}

func getUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	if len(text) < 3 {
		fmt.Println("skipping db insertion")
		return ""
	} else {
		fmt.Println("text looks ok.")
	}
	return text
}
