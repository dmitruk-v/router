package router

import (
	"fmt"
	"log"
	"net/http"
)

type router struct {
	get    *tree
	post   *tree
	put    *tree
	patch  *tree
	delete *tree
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
	}
}

// Select router tree based on http method
func (r *router) selectTree(method string) *tree {
	switch method {
	case http.MethodGet:
		return r.get
	case http.MethodPost:
		return r.post
	case http.MethodPut:
		return r.put
	case http.MethodPatch:
		return r.patch
	case http.MethodDelete:
		return r.delete
	default:
		log.Fatal("Unsupported http method:", method)
		return nil
	}
}

// 1. Select router tree based on http method.
// 2. Iterate through pattern. Breaks pattern string to slashes and words.
// 3. Build new branch in the tree, based on parts extracted from pattern.
// 4. Attach a http handler function to last founded node.
func (r *router) Parse(pattern string, method string, handler http.Handler) *tree {

	child := r.selectTree(method)
	prev := 0
	word := ""

	for i := 0; i < len(pattern); i++ {
		if pattern[i] == '/' {
			// extract word between slashes
			if prev != i {
				word = pattern[prev+1 : i]
				// create word tree node
				child = r.addChild(child, word)
			}
			// create slash tree node
			child = r.addChild(child, "/")
			prev = i
		}
	}
	// last part
	if prev+1 < len(pattern) {
		word := pattern[prev+1:]
		//create word tree node
		child = r.addChild(child, word)
	}

	child.handler = handler
	return child
}

// If parent already contains child with this part, return this existing child.
// Otherwise create a new child node and add it to the parent children list.
func (r *router) addChild(parent *tree, part string) *tree {
	// check parent.children nodes for node with part = word
	child := r.find(parent, part)
	// if founded, return this node
	if child != nil {
		return child
	}
	// otherwise, create new child node
	child = &tree{
		part:     part,
		children: []*tree{},
	}
	parent.children = append(parent.children, child)
	return child
}

// Find child node in parent.children list with provided part.
func (r *router) find(parent *tree, part string) *tree {
	// it is very unlikely, that parent can be nil,
	// but in this case return nil
	if parent == nil {
		return parent
	}
	// If child with equal part already exists, return it.
	for _, chd := range parent.children {
		if chd.part == part {
			return chd
		}
	}
	return nil
}

// Break reqPath into slashes and words. Traverse router tree following
// reqPath. Return the longest matched path end node.
func (r *router) Match(reqPath string, method string) *tree {
	node := r.selectTree(method)
	prev := 0
	word := ""

	for i := 0; i < len(reqPath); i++ {
		if reqPath[i] == '/' {

			if prev != i {
				word = reqPath[prev+1 : i]

				// check word match in tree
				matched := matchPart(node, word)
				if matched == nil {
					// fmt.Printf("MATCH NOT FOUND %v, %v, %v, %v\n", node, word, reqPath, method)
					return node
				}
				node = matched
			}

			// check slash match in tree
			matched := matchPart(node, "/")
			if matched == nil {
				// fmt.Printf("MATCH NOT FOUND %v, %v, %v, %v\n", node, "/", reqPath, method)
				return node
			}
			node = matched

			prev = i
		}
	}
	// get last word
	if prev+1 < len(reqPath) {
		word := reqPath[prev+1:]
		// check last word match in tree
		matched := matchPart(node, word)
		if matched == nil {
			// fmt.Printf("MATCH NOT FOUND %v, %v, %v, %v\n", node, word, reqPath, method)
			return node
		}
		node = matched
	}

	return node
}

func matchPart(parent *tree, part string) *tree {
	for _, ch := range parent.children {
		// if parts have exact match
		if ch.part == part {
			return ch
		}
		// if regex found, and it matches part
		if len(ch.part) > 2 && ch.part[0] == '{' && ch.part[len(ch.part)-1] == '}' {
			return ch
		}
	}
	return nil
}

func (r *router) PrintTree() {
	r.get.Print()
	r.post.Print()
	r.put.Print()
	r.patch.Print()
	r.delete.Print()
}

func (r *router) Add(pattern string, method string, handler http.Handler) {
	node := r.Parse(pattern, method, handler)
	fmt.Println(node)
}
