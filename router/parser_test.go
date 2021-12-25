package router

import "testing"

func TestParser(t *testing.T) {
	table := map[string][]token{
		"/": {
			{tp: SlashType, val: "/"},
		},
		"/users": {
			{tp: SlashType, val: "/"},
			{tp: WordType, val: "users"},
		},
		"/users/": {
			{tp: SlashType, val: "/"},
			{tp: WordType, val: "users"},
			{tp: SlashType, val: "/"},
		},
		"/users/{user_id}": {
			{tp: SlashType, val: "/"},
			{tp: WordType, val: "users"},
			{tp: SlashType, val: "/"},
			{tp: WordType, val: "{user_id}"},
		},
		"//": {
			{tp: SlashType, val: "/"},
			{tp: SlashType, val: "/"},
		},
		"users": {
			{tp: WordType, val: "users"},
		},
		"": {},
	}

	for pattern, tokens := range table {
		p := NewParser(pattern)
		got := p.Parse()
		if !equalsSlice(got, tokens) {
			t.Errorf("got %v, want %v", got, tokens)
		}
	}

}

func equalsSlice(s1, s2 []token) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
