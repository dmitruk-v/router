package main

import (
	"fmt"
	"net/http"

	"github.com/dmitruk-v/router/router"
)

type route struct {
	pattern string
	method  string
	handler http.Handler
}

func main() {

	patterns := []route{
		{pattern: "/users", method: http.MethodGet, handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})},
		{pattern: "/users/{user_id}", method: http.MethodGet, handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})},
		{pattern: "/", method: http.MethodGet, handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})},

		{pattern: "/users/{user_id}", method: http.MethodGet, handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})},
		{pattern: "/users/{user_id}", method: http.MethodPost, handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})},
		{pattern: "/users/{user_id}", method: http.MethodPut, handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})},
		{pattern: "/users/{user_id}", method: http.MethodPatch, handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})},
		{pattern: "/users/{user_id}", method: http.MethodDelete, handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})},

		{pattern: "/users/{user_id}/toys", method: http.MethodGet, handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})},
		{pattern: "/users/{user_id}/toys/{toy_id}", method: http.MethodGet, handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})},
		{pattern: "/users/{user_id}/toys/{toy_id}", method: http.MethodPost, handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})},
		{pattern: "/users/{user_id}/toys/{toy_id}", method: http.MethodPut, handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})},
		{pattern: "/users/{user_id}/toys/{toy_id}", method: http.MethodPatch, handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})},
		{pattern: "/users/{user_id}/toys/{toy_id}", method: http.MethodDelete, handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})},

		{pattern: "/products", method: http.MethodGet, handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})},
		{pattern: "/products/{product_id}/sales", method: http.MethodGet, handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})},
		{pattern: "/products/{product_id}", method: http.MethodGet, handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})},
		{pattern: "/products/{product_id}", method: http.MethodPost, handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})},
		{pattern: "/users/{user_id}/games/{game_id}/", method: http.MethodGet, handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})},
	}
	_ = patterns

	var p1 = "/users/1/games/2/12345/"
	var p2 = "/users/1/"

	r := router.NewRouter()
	for _, ro := range patterns {
		node := r.Parse(ro.pattern, ro.method, ro.handler)
		_ = node
	}

	r.PrintTree()

	node := r.Match(p1, http.MethodGet)
	fmt.Printf("--- FOUNDED MATCH for %s: %v\n", p1, node)

	node = r.Match(p2, http.MethodGet)
	fmt.Printf("--- FOUNDED MATCH for %s: %v\n", p2, node)

	// node2 := r.Match("/users", http.MethodPost)
	// fmt.Println(node2, node2.GetHandler())

	// r.Match("/products/1")
	// parser := router.NewParser(p)
	// tokens := parser.Parse()
	// fmt.Println(tokens)

}
