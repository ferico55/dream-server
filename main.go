package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"server/middleware"
)

const port = "8080"

func handle(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, fmt.Sprintf("Hello from port %q \n", port))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {

}

// Entry point of the package main
func main() {
	// ini dari belakang ke depan
	http.HandleFunc("/", middleware.Chain(rootHandler, middleware.Logging(), middleware.Auth()))
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
