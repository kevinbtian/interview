package main

import (
	"fmt"
	"net/http"
	"github.com/kevinbtian/interview/kong/handler"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, world!")
	})

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, test!")
	})

	h := handler.NewHandler()
	http.HandleFunc("/service", h.GetServicesHandler)

	fmt.Println("Starting server on hocalhost:8080...")
	http.ListenAndServe(":8080", nil)
}
