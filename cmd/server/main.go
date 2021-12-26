package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dmitruk-v/router/router"
)

func main() {
	ro := router.NewRouter()

	ro.HandleFunc("/", http.MethodGet, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello world")
	})

	ro.HandleFunc("/users/{user_id}", http.MethodGet, func(w http.ResponseWriter, r *http.Request) {
		vals := ro.Params(r)
		fmt.Fprintln(w, vals)
	})

	ro.HandleFunc("/users/{user_id}/games/{game_id}", http.MethodGet, func(w http.ResponseWriter, r *http.Request) {
		vals := ro.Params(r)
		fmt.Fprintln(w, vals)
	})

	log.Println("Server started at localhost:4000")
	if err := http.ListenAndServe("localhost:4000", ro); err != nil {
		log.Fatal(err)
	}
}
