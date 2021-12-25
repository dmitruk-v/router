package router

import (
	"fmt"
	"net/http"
	"strings"
)

type tree struct {
	part     string
	method   string
	handler  http.Handler
	children []*tree
}

// Deep compare two trees
func (t *tree) Equals(t2 *tree) bool {
	if t.part != t2.part {
		return false
	}
	if t.method != t2.method {
		return false
	}
	if t.handler != t2.handler {
		return false
	}
	if len(t.children) != len(t2.children) {
		return false
	}
	for i := 0; i < len(t.children); i++ {
		eq := t.children[i].Equals(t2.children[i])
		if !eq {
			return false
		}
	}
	return true
}

// Pretty print a tree
func (t *tree) Print() {
	var sb = strings.Builder{}

	if t == nil {
		fmt.Println(nil)
	}

	indent := func(depth int, spacer string) string {
		sb := strings.Builder{}
		for i := 0; i < depth; i++ {
			sb.WriteString(spacer)
		}
		return sb.String()
	}

	var fn func(t *tree, depth int)
	fn = func(t *tree, depth int) {
		sb.WriteString(fmt.Sprintf("%s%s %v\n", indent(depth, "- "), t.part, t.handler))
		for _, node := range t.children {
			fn(node, depth+1)
		}
	}
	fn(t, 0)
	fmt.Print(sb.String())
}

func (t *tree) GetHandler() http.Handler {
	return t.handler
}
