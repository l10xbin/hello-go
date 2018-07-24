package main

import (
	"fmt"
	"net/http"
	"os"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func sayPongJSON(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, `{"message":"pong"}`)
}

func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/ping", sayPongJSON)

	// get port env var
	port := "8080"
	portEnv := os.Getenv("PORT")
	if len(portEnv) > 0 {
		port = portEnv
	}
	fmt.Printf("starting up on port: %s ...", port)

	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	// listen and serve on 0.0.0.0:8080 by default
	// set environment variable PORT if you want to change port
}