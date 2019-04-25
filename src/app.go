package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/l10xbin/hello-go/src/handler"
)

func main() {
	http.HandleFunc("/", handler.SayHello)
	http.HandleFunc("/ping", handler.SayPongJSON)

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
