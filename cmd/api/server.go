package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from root route"))
}

func teachersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Println(r.URL.Path)
		userID := strings.TrimSuffix(
			strings.TrimPrefix(r.URL.Path, "/teachers/"),
			"/",
		)

		name := r.URL.Query().Get("name")
		if name == "" {
			name = "Emil"
		}

		fmt.Println(name)
		fmt.Println(userID)

		w.Write([]byte("Hello from teachers route - GET"))
	case http.MethodPost:
		w.Write([]byte("Hello from teachers route - POST"))
	case http.MethodPut:
		w.Write([]byte("Hello from teachers route - PUT"))
	case http.MethodPatch:
		w.Write([]byte("Hello from teachers route - PATCH"))
	case http.MethodDelete:
		w.Write([]byte("Hello from teachers route - DELETE"))
	}
}

func studentsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello from students route - GET"))
	case http.MethodPost:
		w.Write([]byte("Hello from students route - POST"))
	case http.MethodPut:
		w.Write([]byte("Hello from students route - PUT"))
	case http.MethodPatch:
		w.Write([]byte("Hello from students route - PATCH"))
	case http.MethodDelete:
		w.Write([]byte("Hello from students route - DELETE"))
	}
}

func execsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello from execs route - GET"))
	case http.MethodPost:
		w.Write([]byte("Hello from execs route - POST"))
	case http.MethodPut:
		w.Write([]byte("Hello from execs route - PUT"))
	case http.MethodPatch:
		w.Write([]byte("Hello from execs route - PATCH"))
	case http.MethodDelete:
		w.Write([]byte("Hello from execs route - DELETE"))
	}
}

func main() {
	port := ":3000"

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/teachers/", teachersHandler)
	http.HandleFunc("/students/", studentsHandler)
	http.HandleFunc("/execs/", execsHandler)

	fmt.Println("Server is running on port:", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalln("Error starting the server", err)
	}
}
