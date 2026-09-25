package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	"api_proj/internal/api/middlewares"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from root route"))
}

func teachersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
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
	cert := "cert.pem"
	key := "key.pem"

	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/teachers/", teachersHandler)
	mux.HandleFunc("/students/", studentsHandler)
	mux.HandleFunc("/execs/", execsHandler)

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	server := &http.Server{
		Addr:      port,
		Handler:   middlewares.SecurityHeaders(middlewares.Cors(mux)),
		TLSConfig: tlsConfig,
	}

	fmt.Println("Server is running on port", port)
	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Error starting the server", err)
	}
}
