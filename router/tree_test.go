package router

import (
	"fmt"
	"net/http"
	"testing"
)

func TestEquals(t *testing.T) {
	t1 := &tree{
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
			{part: "products", method: http.MethodGet, handler: nil, children: []*tree{
				{part: "/", method: http.MethodGet, handler: nil, children: []*tree{
					{part: "{product_id}", method: http.MethodGet, handler: nil, children: []*tree{}},
				}},
			}},
		}}

	t2 := &tree{
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
			{part: "products", method: http.MethodGet, handler: nil, children: []*tree{
				{part: "/", method: http.MethodGet, handler: nil, children: []*tree{
					{part: "{product_id}", method: http.MethodGet, handler: nil, children: []*tree{}},
				}},
			}},
		}}

	if !t1.Equals(t2) {
		t.Errorf("trees are not equal")
		fmt.Printf("\n*** t1 ***************************\n\n")
		t1.Print()
		fmt.Printf("\n*** NOT EQUALS t2 ****************\n\n")
		t2.Print()
		fmt.Printf("\n")
	}
}
