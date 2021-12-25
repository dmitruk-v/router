package router

import (
	"strings"
)

type tokenType string

const (
	SlashType tokenType = "SLASH"
	WordType  tokenType = "WORD"
)

var SlashToken = token{tp: SlashType, val: "/"}

type token struct {
	tp  tokenType
	val string
}

type parser struct {
	pattern string
	index   int
	tokens  []token
}

func NewParser(pattern string) *parser {
	return &parser{
		pattern: pattern,
		index:   0,
		tokens:  []token{},
	}
}

// Break pattern string to the tokens list
func (p *parser) Parse() []token {
	for {
		if p.eol() {
			break
		}
		curr := p.pattern[p.index]
		if curr == '/' {
			SlashType := p.slash()
			p.tokens = append(p.tokens, SlashType)
			continue
		}
		WordType := p.word()
		p.tokens = append(p.tokens, WordType)
	}
	return p.tokens
}

// End of line was reached
func (p *parser) eol() bool {
	return p.index >= len(p.pattern)
}

// Consume a slash
func (p *parser) slash() token {
	p.index++
	// Slash tokens always the same, so return same token
	return SlashToken
}

// Consume a word, letters between slashes.
func (p *parser) word() token {
	b := strings.Builder{}
	for {
		if p.eol() || p.pattern[p.index] == '/' {
			break
		}
		b.WriteByte(p.pattern[p.index])
		p.index++
	}
	return token{tp: WordType, val: b.String()}
}
