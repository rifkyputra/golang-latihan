package logrequest

import (
	"errors"
	apphttp "examples/coba/internal/http"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"
)

func RunApp() {
	http.HandleFunc("/", LogRequestToFile(apphttp.Greet))
	http.HandleFunc("/hello", LogRequestToFile(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello There")
	}))

	apphttp.Listen(":8080").Serve()
}

func LogRequestToFile(handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	// Log the request to a file
	f, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		ip, _ := getIPAddr(r)
		f.WriteString(r.URL.Path + "----" + time.Now().String() + " " + ip + "\n")
		handler(w, r)
		defer f.Close()
	}
}

func getIPAddr(r *http.Request) (string, error) {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", err
	}

	userIP := net.ParseIP(ip)
	if userIP == nil {
		return "", errors.New("user IP is nil")
	}

	return userIP.String(), nil
}
