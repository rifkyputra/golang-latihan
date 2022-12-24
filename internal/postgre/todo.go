package postgre

import (
	"database/sql"
	"fmt"
)

type Todo struct {
	ID    int
	Title string
	Check bool
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

func UpdateTodo(db *sql.DB, todo Todo) int {
	updateStmt := fmt.Sprintf(`update todo set title = '%s', "check" = %t where id = %d;`, todo.Title, todo.Check, todo.ID)
	fmt.Println(updateStmt)

	_, err := db.Exec(updateStmt)

	if err != nil {
		fmt.Println(err)
		return 0

	}

	return 1

}
