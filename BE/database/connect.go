package database

import (
	"Portfolio/ent"
	"fmt"
	"log"

	"entgo.io/ent/dialect"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 1234
	user     = "kate"
	password = "secret"
	dbname   = "Assignment4"
)

var Client *ent.Client

func Connect() {
	connInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := ent.Open(dialect.Postgres, connInfo)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to database!")

	Client = db
}
