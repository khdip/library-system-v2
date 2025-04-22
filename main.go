package main

import (
	"fmt"
	"log"
	"net/http"

	"practice/library-system-v2/handler"

	"github.com/gorilla/schema"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	var createTable = `
			CREATE TABLE IF NOT EXISTS books (
				id serial,
				book_name text,
				author text,
				category text,
				book_description text,
				book_cover text,
				is_available boolean,

				primary key(id)
			);`

	db, err := sqlx.Connect("postgres", "user=postgres password=Anubis0912 dbname=library_v2 sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	db.MustExec(createTable)

	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)
	r := handler.GetHandler(db, decoder)

	fmt.Println("Server Starting...")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("Server Not Found", err)
	}
}
