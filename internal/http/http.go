package apphttp

import (
	"fmt"
	"net/http"
	"time"
)

func (l *ListenAndServe) Serve() {
	http.ListenAndServe(l.port, nil)
}

func Listen(port string) *ListenAndServe {
	return &ListenAndServe{port: port}
}

func Greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

type ListenAndServe struct {
	port string
	w    http.ResponseWriter
	r    *http.Request
}
