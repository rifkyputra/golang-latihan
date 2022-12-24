package postgre

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

func ConnectToPostgresql() *sql.DB {
	cfg := "postgres://apps:apps@localhost:5432/todo?sslmode=disable"
	db, err := sql.Open("postgres", cfg)

	if err != nil {
		fmt.Println("Fail ")
		panic(err)
	}

	if err = db.Ping(); err != nil {
		fmt.Println("No Response")
		panic(err)
	}

	return db
}

type Handler = func(w http.ResponseWriter, r *http.Request)
