package postgre

import (
	"database/sql"
	"encoding/json"
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

func CreateTodoTable(db *sql.DB) {
	createStmt := `CREATE TABLE IF NOT EXISTS "todo" (
		"id" SERIAL PRIMARY KEY,
		"title" VARCHAR(100) NOT NULL,
		"check" BOOLEAN NOT NULL
	);`

	_, err := db.Exec(createStmt)
	if err != nil {
		fmt.Println(err)
	}
}

func InsertTodo(title string, db *sql.DB) int {
	insertStmt := fmt.Sprintf(`insert into "todo"("title", "check") values('%s', false)`, title)
	_, e := db.Exec(insertStmt)

	if e != nil {
		fmt.Println(e)
		return 0
	}

	return 1
}

type Todo struct {
	ID    int
	Title string
	Check bool
}

func getTodos(db *sql.DB) []Todo {
	rows, err := db.Query("SELECT * FROM todo")
	if err != nil {
		fmt.Println(err)
	}

	var todos []Todo

	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID, &todo.Title, &todo.Check)
		if err != nil {
			fmt.Println(err)
		}
		todos = append(todos, todo)
	}

	return todos
}

type Handler = func(w http.ResponseWriter, r *http.Request)

func HandleTodo(db *sql.DB) {
	http.HandleFunc("/todo", HandleGetTodo(db, HandlePostTodo(db, nil)))
}

func HandleGetTodo(db *sql.DB, in Handler) Handler {

	out := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			todos := getTodos(db)
			json.NewEncoder(w).Encode(todos)
		}

		if in != nil {
			in(w, r)
		}
	}

	return out

}

func HandlePostTodo(db *sql.DB, in Handler) Handler {
	out := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			var todo Todo
			if r.Body == nil {
				http.Error(w, "Please send a request body", 400)
				return
			}
			err := json.NewDecoder(r.Body).Decode(&todo)
			if err != nil {
				http.Error(w, err.Error(), 400)
				return
			}

			if todo.Title == "" {
				http.Error(w, "Title must not empty", 400)
				return
			}

			InsertTodo(todo.Title, db)

			fmt.Fprintf(w, "Success")
		}

		if in != nil {
			in(w, r)
		}
	}

	return out
}
