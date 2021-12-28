package router

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

type branchTest struct {
	pattern string
	handler http.Handler
	node    *tree
}

func newBranchTest(pattern string, handler http.Handler, node *tree) branchTest {
	bt := branchTest{
		pattern: pattern,
		handler: handler,
		node:    node,
	}
	bt.node.handler = handler
	return bt
}

var handlers = []http.Handler{
	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
}

var table = []branchTest{
	newBranchTest("/", handlers[0], &tree{part: "/", children: []*tree{}}),
	newBranchTest("/users", handlers[1], &tree{part: "users", children: []*tree{}}),
	newBranchTest("/users/", handlers[2], &tree{part: "/", children: []*tree{}}),
	newBranchTest("/users/{user_id}", handlers[3], &tree{part: "{user_id}", children: []*tree{}}),
	newBranchTest("/users/{user_id}/games/{game_id}", handlers[4], &tree{part: "{game_id}", children: []*tree{}}),
	newBranchTest("/users/{user_id}/toys/{toy_id}", handlers[5], &tree{part: "{toy_id}", children: []*tree{}}),
	newBranchTest("/users/{user_id}/toys", handlers[6], &tree{part: "toys", children: []*tree{}}),
	newBranchTest("/users/{user_id}/books/{book_id}", handlers[7], &tree{part: "{book_id}", children: []*tree{}}),
	newBranchTest("/products", handlers[8], &tree{part: "products", children: []*tree{}}),
	newBranchTest("/products/{product_id}", handlers[9], &tree{part: "{product_id}", children: []*tree{}}),
}

func TestAddChild(t *testing.T) {
	router := NewRouter()
	child := addChild(router.get, "/")
	want := &tree{part: "/", children: []*tree{}}
	if !child.Equals(want) {
		t.Errorf("got %v, want %v", child, want)
	}
}

func TestParse(t *testing.T) {

	want := &tree{part: "GET", children: []*tree{
		{part: "/", handler: handlers[0], children: []*tree{
			{part: "users", handler: handlers[1], children: []*tree{
				{part: "/", handler: handlers[2], children: []*tree{
					{part: "{user_id}", handler: handlers[3], children: []*tree{
						{part: "/", handler: nil, children: []*tree{
							{part: "games", handler: nil, children: []*tree{
								{part: "/", handler: nil, children: []*tree{
									{part: "{game_id}", handler: handlers[4], children: []*tree{}},
								}},
							}},
							{part: "toys", handler: handlers[6], children: []*tree{
								{part: "/", handler: nil, children: []*tree{
									{part: "{toy_id}", handler: handlers[5], children: []*tree{}},
								}},
							}},
							{part: "books", handler: nil, children: []*tree{
								{part: "/", handler: nil, children: []*tree{
									{part: "{book_id}", handler: handlers[7], children: []*tree{}},
								}},
							}},
						}},
					}},
				}},
			}},
			{part: "products", handler: handlers[8], children: []*tree{
				{part: "/", handler: nil, children: []*tree{
					{part: "{product_id}", handler: handlers[9], children: []*tree{}},
				}},
			}},
		}},
	}}

	r := NewRouter()
	for _, bt := range table {
		node := r.parse(bt.pattern, http.MethodGet, bt.handler)
		if node.part != bt.node.part || reflect.ValueOf(node.handler) != reflect.ValueOf(bt.handler) {
			t.Errorf("nodes are not equal, got %v, want %v", node, bt.node)
			fmt.Printf("\n*** t1 ***************************\n\n")
			node.Print()
			fmt.Printf("\n*** NOT EQUALS t2 ****************\n\n")
			bt.node.Print()
			fmt.Printf("\n")
		}
	}
	if !r.get.Equals(want) {
		t.Errorf("trees are not equal, got %v, want %v", r.get, want)
		fmt.Printf("\n*** t1 ***************************\n\n")
		r.get.Print()
		fmt.Printf("\n*** NOT EQUALS t2 ****************\n\n")
		want.Print()
		fmt.Printf("\n")
	}
}

type matchTest struct {
	pattern string
	method  string
	handler http.Handler
}

func BenchmarkMatch(b *testing.B) {
	table := []matchTest{
		{pattern: "/users", method: http.MethodGet, handler: nil},
		{pattern: "/users/{user_id}/games/{game_id}", method: http.MethodGet, handler: nil},
		{pattern: "/users/{user_id}/games", method: http.MethodGet, handler: nil},
		{pattern: "/users/{user_id}", method: http.MethodGet, handler: nil},
		{pattern: "/users/{user_id}/toys", method: http.MethodGet, handler: nil},
		{pattern: "/users/{user_id}/toys/{toy_id}", method: http.MethodGet, handler: nil},
		{pattern: "/products", method: http.MethodGet, handler: nil},
		{pattern: "/products/{product_id}/sales", method: http.MethodGet, handler: nil},
		{pattern: "/products/{product_id}", method: http.MethodGet, handler: nil},
		{pattern: "/users/{user_id}/games/{game_id}/", method: http.MethodGet, handler: nil},
	}
	r := NewRouter()
	for _, mt := range table {
		r.parse(mt.pattern, mt.method, mt.handler)
	}

	for i := 0; i < b.N; i++ {
		ctx := context.WithValue(context.Background(), routeParamsKey, make(map[string]string))
		node := r.match(ctx, "/users/123/games/456/bla", http.MethodGet)
		_ = node
	}
}
