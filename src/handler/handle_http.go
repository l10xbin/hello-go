package handler

import (
	"fmt"
	"net/http"
	"os"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	hostName, err := os.Hostname()
	hello := "Hello world"
	if err == nil {
		hello = "Hello World from " + hostName + " !"
	}

	fmt.Fprint(w, hello)
}

func SayPongJSON(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, `{"message":"pong"}`)
}
