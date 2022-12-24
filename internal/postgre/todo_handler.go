package postgre

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleTodo(db *sql.DB) {
	http.HandleFunc("/todo", HandleGetTodo(db, HandlePostTodo(db, HandlePutTodo(db, nil))))
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

func HandlePutTodo(db *sql.DB, in Handler) Handler {
	out := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "PUT" {
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

			query := UpdateTodo(db, todo)

			if query == 0 {
				http.Error(w, "Todo not found", 400)
			} else {

				fmt.Fprintf(w, "Success")
			}

		}

		if in != nil {
			in(w, r)
		}
	}

	return out

}
