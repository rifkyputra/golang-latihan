package main

import (
	serveAndLog "examples/coba/internal/log_request"
	db "examples/coba/internal/postgre"
	"fmt"
)

func main() {
	fmt.Println("Starting...")
	dbConn := db.ConnectToPostgresql()

	fmt.Println("Connected To Database.")

	db.HandleTodo(dbConn)
	db.CreateTodoTable(dbConn)
	// db.InsertTodo("DR", dbConn)

	serveAndLog.RunApp()

	fmt.Println("Server Running")

}
