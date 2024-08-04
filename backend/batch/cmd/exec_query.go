package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=postgres dbname=knowledge sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	query, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(string(query))
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully executed query from file:", os.Args[1])
}
