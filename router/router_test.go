package router

import (
	"net/http"
	"testing"
)

type branchTest struct {
	pattern string
	root    *tree
	last    *tree
}

var table = []branchTest{
	{
		pattern: "/",
		root:    &tree{part: "/", method: http.MethodGet, handler: nil, children: []*tree{}},
		last:    &tree{part: "/", method: http.MethodGet, handler: nil, children: []*tree{}},
	},
	{
		pattern: "/users/",
		root: &tree{
			part: "/", method: http.MethodGet, handler: nil, children: []*tree{
				{part: "users", method: http.MethodGet, handler: nil, children: []*tree{
					{part: "/", method: http.MethodGet, handler: nil, children: []*tree{}},
				}},
			}},
		last: &tree{part: "/", method: http.MethodGet, handler: nil, children: []*tree{}},
	},
	{
		pattern: "/users/{user_id}/games/{game_id}",
		root: &tree{
			part: "/", method: http.MethodGet, handler: nil, children: []*tree{
				{part: "users", method: http.MethodGet, handler: nil, children: []*tree{
					{part: "/", method: http.MethodGet, handler: nil, children: []*tree{
						{part: "{user_id}", method: http.MethodGet, handler: nil, children: []*tree{
							{part: "/", method: http.MethodGet, handler: nil, children: []*tree{
								{part: "games", method: http.MethodGet, handler: nil, children: []*tree{
									{part: "/", method: http.MethodGet, handler: nil, children: []*tree{
										{part: "{game_id}", method: http.MethodGet, handler: nil, children: []*tree{}},
									}},
								}},
							}},
						}},
					}},
				}},
			}},
		last: &tree{part: "{game_id}", method: http.MethodGet, handler: nil, children: []*tree{}},
	},
	{
		pattern: "/users/{user_id}/toys/{toy_id}",
		root: &tree{
			part: "/", method: http.MethodGet, handler: nil, children: []*tree{
				{part: "users", method: http.MethodGet, handler: nil, children: []*tree{
					{part: "/", method: http.MethodGet, handler: nil, children: []*tree{
						{part: "{user_id}", method: http.MethodGet, handler: nil, children: []*tree{
							{part: "/", method: http.MethodGet, handler: nil, children: []*tree{
								{part: "games", method: http.MethodGet, handler: nil, children: []*tree{
									{part: "/", method: http.MethodGet, handler: nil, children: []*tree{
										{part: "{game_id}", method: http.MethodGet, handler: nil, children: []*tree{}},
									}},
								}},
								{part: "toys", method: http.MethodGet, handler: nil, children: []*tree{
									{part: "/", method: http.MethodGet, handler: nil, children: []*tree{
										{part: "{toy_id}", method: http.MethodGet, handler: nil, children: []*tree{}},
									}},
								}},
							}},
						}},
					}},
				}},
			}},
		last: &tree{part: "{toy_id}", method: http.MethodGet, handler: nil, children: []*tree{}},
	},
	{
		pattern: "/users/{user_id}/books/{book_id}",
		root: &tree{
			part: "/", method: http.MethodGet, handler: nil, children: []*tree{
				{part: "users", method: http.MethodGet, handler: nil, children: []*tree{
					{part: "/", method: http.MethodGet, handler: nil, children: []*tree{
						{part: "{user_id}", method: http.MethodGet, handler: nil, children: []*tree{
							{part: "/", method: http.MethodGet, handler: nil, children: []*tree{
								{part: "games", method: http.MethodGet, handler: nil, children: []*tree{
									{part: "/", method: http.MethodGet, handler: nil, children: []*tree{
										{part: "{game_id}", method: http.MethodGet, handler: nil, children: []*tree{}},
									}},
								}},
								{part: "toys", method: http.MethodGet, handler: nil, children: []*tree{
									{part: "/", method: http.MethodGet, handler: nil, children: []*tree{
										{part: "{toy_id}", method: http.MethodGet, handler: nil, children: []*tree{}},
									}},
								}},
								{part: "books", method: http.MethodGet, handler: nil, children: []*tree{
									{part: "/", method: http.MethodGet, handler: nil, children: []*tree{
										{part: "{book_id}", method: http.MethodGet, handler: nil, children: []*tree{}},
									}},
								}},
							}},
						}},
					}},
				}},
			}},
		last: &tree{part: "{book_id}", method: http.MethodGet, handler: nil, children: []*tree{}},
	},
	{
		pattern: "/products/{product_id}",
		root: &tree{
			part: "/", method: http.MethodGet, handler: nil, children: []*tree{
				{part: "users", method: http.MethodGet, handler: nil, children: []*tree{
					{part: "/", method: http.MethodGet, handler: nil, children: []*tree{
						{part: "{user_id}", method: http.MethodGet, handler: nil, children: []*tree{
							{part: "/", method: http.MethodGet, handler: nil, children: []*tree{
								{part: "games", method: http.MethodGet, handler: nil, children: []*tree{
									{part: "/", method: http.MethodGet, handler: nil, children: []*tree{
										{part: "{game_id}", method: http.MethodGet, handler: nil, children: []*tree{}},
									}},
								}},
								{part: "toys", method: http.MethodGet, handler: nil, children: []*tree{
									{part: "/", method: http.MethodGet, handler: nil, children: []*tree{
										{part: "{toy_id}", method: http.MethodGet, handler: nil, children: []*tree{}},
									}},
								}},
								{part: "books", method: http.MethodGet, handler: nil, children: []*tree{
									{part: "/", method: http.MethodGet, handler: nil, children: []*tree{
										{part: "{book_id}", method: http.MethodGet, handler: nil, children: []*tree{}},
									}},
								}},
							}},
						}},
					}},
				}},
				{part: "products", method: http.MethodGet, handler: nil, children: []*tree{
					{part: "/", method: http.MethodGet, handler: nil, children: []*tree{
						{part: "{product_id}", method: http.MethodGet, handler: nil, children: []*tree{}},
					}},
				}},
			}},
		last: &tree{part: "{product_id}", method: http.MethodGet, handler: nil, children: []*tree{}},
	},
}

func TestAddChild(t *testing.T) {
	router := NewRouter()
	child := router.addChild(router.get, "/")
	want := &tree{part: "/", children: []*tree{}}
	if !child.Equals(want) {
		t.Errorf("got %v, want %v", child, want)
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
		r.Parse(mt.pattern, mt.method, mt.handler)
	}

	for i := 0; i < b.N; i++ {
		node := r.Match("/users/12345/games/12345/", http.MethodGet)
		_ = node
	}
}
