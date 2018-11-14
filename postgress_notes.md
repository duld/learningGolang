# Postgressql notes

[command cheat sheet](http://www.postgresqltutorial.com/postgresql-cheat-sheet/)
[go and postgres tut](https://www.calhoun.io/querying-for-multiple-records-with-gos-sql-package/)

## Useful Commands for psql.exe

\list (\l) - list all defined databases.
\connect < dbname >(\c < dbname >) - connect to a different database.
\dt - list the tables under the database.
\du or \du+ - list all user accounts

## Connecting to our Database

Important info.

* Postgress install path: Program Files\PostgreSQL\11\bin\pysql.exe
* Golang postgresql drivergithub.com/lib/pq

Steps to connect

1. import: __database/sql__
2. install: __github.com/lib/pq__
3. Use a connection string that adheres to this format: "host=hostname port=portNum user=username password=passVal dbname=dbToConnectTo sslmode=disable"
4. Open a connection to the database using __sql.Open()__
5. use __db.Ping()__ to establish a live connection.

Example code

```Golang

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
}

```

## pgadmin4 notes

* We cannot save or edit data in a table if we are missing a primary key.