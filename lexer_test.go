package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lexSyntax(t *testing.T) {
	tests := []struct {
		src        string
		cursor     int
		wantValue  string
		wantCursor int
	}{
		{"lol )", 4, ")", 5},
		{"lm(ao", 2, "(", 3},
	}

	for _, test := range tests {
		cursor, token := lexSyntax([]rune(test.src), test.cursor)
		assert.Equal(t, test.wantCursor, cursor)
		assert.Equal(t, test.wantValue, token.value)
		assert.Equal(t, syntaxToken, token.kind)
	}
}

func Test_lexInteger(t *testing.T) {
	tests := []struct {
		src        string
		cursor     int
		wantValue  string
		wantCursor int
	}{
		{"foo 123", 4, "123", 7},
		{"foo 12 3", 4, "12", 6},
		{"foo 12a 3", 4, "12", 6},
	}

	for _, test := range tests {
		cursor, token := lexInteger([]rune(test.src), test.cursor)
		assert.Equal(t, test.wantCursor, cursor)
		assert.Equal(t, test.wantValue, token.value)
		assert.Equal(t, integerToken, token.kind)
	}
}

func Test_lexIdentifier(t *testing.T) {
	tests := []struct {
		src        string
		cursor     int
		wantValue  string
		wantCursor int
	}{
		{"123 ab + ", 4, "ab", 6},
		{"123 ab123 + ", 4, "ab123", 9},
	}

	for _, test := range tests {
		cursor, token := lexIdentifier([]rune(test.src), test.cursor)
		assert.Equal(t, test.wantCursor, cursor)
		assert.Equal(t, test.wantValue, token.value)
		assert.Equal(t, identifierToken, token.kind)
	}
}

func Test_lex(t *testing.T) {
	tests := []struct {
		src        string
		wantTokens []token
	}{
		{
			" ( + 13 2  )",
			[]token{
				{
					value:    "(",
					kind:     syntaxToken,
					location: 1,
				},
				{
					value:    "+",
					kind:     identifierToken,
					location: 3,
				},
				{
					value:    "13",
					kind:     integerToken,
					location: 5,
				},
				{
					value:    "2",
					kind:     integerToken,
					location: 8,
				},
				{
					value:    ")",
					kind:     syntaxToken,
					location: 11,
				},
			},
		},
	}

	for _, test := range tests {
		tokens := lex(test.src)
		assert.Equal(t, test.wantTokens, tokens)
	}
}
