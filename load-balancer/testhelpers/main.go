package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	Run(os.Getenv("PORT"))
}

func Run(port string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong from %s", port)
		fmt.Println("pong")
	})
	http.ListenAndServe(port, nil)
}
