package v1

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

type routeKey int

const routeParamsKey routeKey = 1

type router struct {
	get     *tree
	post    *tree
	put     *tree
	patch   *tree
	delete  *tree
	connect *tree
	head    *tree
	options *tree
}

func NewRouter() *router {
	return &router{
		get: &tree{
			part:     "GET",
			children: []*tree{},
		},
		post: &tree{
			part:     "POST",
			children: []*tree{},
		},
		put: &tree{
			part:     "PUT",
			children: []*tree{},
		},
		patch: &tree{
			part:     "PATCH",
			children: []*tree{},
		},
		delete: &tree{
			part:     "DELETE",
			children: []*tree{},
		},
		connect: &tree{
			part:     "CONNECT",
			children: []*tree{},
		},
		head: &tree{
			part:     "HEAD",
			children: []*tree{},
		},
		options: &tree{
			part:     "OPTIONS",
			children: []*tree{},
		},
	}
}

// Add http.Handler for the given pattern and method
func (ro *router) Handle(pattern string, method string, handler http.Handler) {
	ro.parse(pattern, method, handler)
}

// Add http.HandlerFunc for the given pattern and method
func (ro *router) HandleFunc(pattern string, method string, handler http.HandlerFunc) {
	ro.parse(pattern, method, handler)
}

// Get url params from context
func (ro *router) Params(r *http.Request) map[string]string {
	m, ok := r.Context().Value(routeParamsKey).(map[string]string)
	if !ok {
		panic("router params map not found in context")
	}
	return m
}

// Select router tree based on http method
func (ro *router) selectTree(method string) *tree {
	switch method {
	case http.MethodGet:
		return ro.get
	case http.MethodPost:
		return ro.post
	case http.MethodPut:
		return ro.put
	case http.MethodPatch:
		return ro.patch
	case http.MethodDelete:
		return ro.delete
	case http.MethodConnect:
		return ro.connect
	case http.MethodHead:
		return ro.head
	case http.MethodOptions:
		return ro.options
	default:
		panic(fmt.Sprintf("Unsupported http method: %s", method))
	}
}

// Build new branch in the tree, based on parts extracted from pattern.
// Attach a http handler function to last founded node.
func (ro *router) parse(pattern string, method string, handler http.Handler) *tree {

	child := ro.selectTree(method)
	prev := 0
	word := ""

	for i := 0; i < len(pattern); i++ {
		if pattern[i] == '/' {
			// extract word between slashes
			if prev != i {
				word = pattern[prev+1 : i]
				// create word tree node, or select existing
				child = addChild(child, word)
			}
			// create slash tree node, or select existing
			child = addChild(child, "/")
			prev = i
		}
	}
	// last part
	if prev+1 < len(pattern) {
		word := pattern[prev+1:]
		//create word tree node, or select existing
		child = addChild(child, word)
	}

	// attach handler to the last created child node
	child.handler = handler
	return child
}

// If parent already contains child with this part, return this existing child.
// Otherwise create a new child node and add it to the parent children list.
func addChild(parent *tree, part string) *tree {
	// check parent.children nodes for node with equal part and return it
	for _, child := range parent.children {
		if child.part == part {
			return child
		}
	}
	// otherwise, create new child node
	child := &tree{
		part:     part,
		children: []*tree{},
	}
	parent.children = append(parent.children, child)
	return child
}

// Break reqPath into slashes and words. Traverse router tree following
// reqPath. Return the farest matched node.
func (ro *router) match(ctx context.Context, reqPath string, method string) *tree {

	node := ro.selectTree(method)
	prev := 0
	word := ""

	// let's say, that max number of url parts would be 32
	parts := [32]string{}
	partIndex := 0

	for i := 0; i < len(reqPath); i++ {
		if reqPath[i] == '/' {
			if prev != i {
				// word
				word = reqPath[prev+1 : i]
				parts[partIndex] = word
				partIndex++
			}
			// slash
			parts[partIndex] = "/"
			partIndex++
			prev = i
		}
	}
	// last word
	if prev+1 < len(reqPath) {
		word := reqPath[prev+1:]
		parts[partIndex] = word
		partIndex++
	}

	// sequentially match collected part to node of the tree,
	// and going deep through the tree
	for _, part := range parts {
		matched := matchPart(ctx, node, part)
		if matched == nil {
			return node
		}
		node = matched
	}

	return node
}

// Match parent children nodes, and find a child that match a part.
// If nothing found, return nil.
func matchPart(ctx context.Context, parent *tree, part string) *tree {
	// it is very unlikely, that parent can be nil,
	// but in this case return nil
	if parent == nil {
		return parent
	}
	for _, child := range parent.children {
		// if parts have exact match
		if child.part == part {
			return child
		}
		// if regex found, and it matches part
		if len(child.part) > 2 && child.part[0] == '{' && child.part[len(child.part)-1] == '}' {
			key := child.part[1 : len(child.part)-1]
			m, ok := ctx.Value(routeParamsKey).(map[string]string)
			if !ok {
				log.Fatal("router params map not found in context")
			}
			m[key] = part
			return child
		}
	}
	return nil
}

// http.Handler implementation
func (ro *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctx := context.WithValue(req.Context(), routeParamsKey, make(map[string]string))
	node := ro.match(ctx, req.URL.Path, req.Method)

	// we do not have handler for this route
	if node.handler == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	node.handler.ServeHTTP(w, req.WithContext(ctx))
}

// Print whole router tree
func (ro *router) PrintTree() {
	ro.get.Print()
	ro.post.Print()
	ro.put.Print()
	ro.patch.Print()
	ro.delete.Print()
	ro.connect.Print()
	ro.head.Print()
	ro.options.Print()
}
