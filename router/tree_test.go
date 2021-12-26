package router

import (
	"fmt"
	"testing"
)

func TestEquals(t *testing.T) {
	t1 := &tree{
		part: "/", handler: nil, children: []*tree{
			{part: "users", handler: nil, children: []*tree{
				{part: "/", handler: nil, children: []*tree{
					{part: "{user_id}", handler: nil, children: []*tree{
						{part: "/", handler: nil, children: []*tree{
							{part: "games", handler: nil, children: []*tree{
								{part: "/", handler: nil, children: []*tree{
									{part: "{game_id}", handler: nil, children: []*tree{}},
								}},
							}},
							{part: "toys", handler: nil, children: []*tree{
								{part: "/", handler: nil, children: []*tree{
									{part: "{toy_id}", handler: nil, children: []*tree{}},
								}},
							}},
						}},
					}},
				}},
			}},
			{part: "products", handler: nil, children: []*tree{
				{part: "/", handler: nil, children: []*tree{
					{part: "{product_id}", handler: nil, children: []*tree{}},
				}},
			}},
		}}

	t2 := &tree{
		part: "/", handler: nil, children: []*tree{
			{part: "users", handler: nil, children: []*tree{
				{part: "/", handler: nil, children: []*tree{
					{part: "{user_id}", handler: nil, children: []*tree{
						{part: "/", handler: nil, children: []*tree{
							{part: "games", handler: nil, children: []*tree{
								{part: "/", handler: nil, children: []*tree{
									{part: "{game_id}", handler: nil, children: []*tree{}},
								}},
							}},
							{part: "toys", handler: nil, children: []*tree{
								{part: "/", handler: nil, children: []*tree{
									{part: "{toy_id}", handler: nil, children: []*tree{}},
								}},
							}},
						}},
					}},
				}},
			}},
			{part: "products", handler: nil, children: []*tree{
				{part: "/", handler: nil, children: []*tree{
					{part: "{product_id}", handler: nil, children: []*tree{}},
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
