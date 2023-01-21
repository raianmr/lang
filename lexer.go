package main

import "unicode"

type tokenKind uint

const (
	syntax tokenKind = iota
	integer
	identifier
)

type token struct {
	value    string
	kind     tokenKind
	location int
}

func (t token) debug(src []rune) {}

func lex(raw string) []token {
	src := []rune(raw)
	var tokens []token

	var t *token
	cursor := 0
	for cursor < len(src) {
		cursor = trimSpaces(src, cursor)

		if cursor == len(src) {
			break
		}

		if cursor, t = lexSyntax(src, cursor); t != nil {
			tokens = append(tokens, *t)
			continue
		}

		if cursor, t = lexInteger(src, cursor); t != nil {
			tokens = append(tokens, *t)
			continue
		}

		if cursor, t = lexIdentifier(src, cursor); t != nil {
			tokens = append(tokens, *t)
			continue
		}

		panic("cant lex")
	}

	return tokens
}

func trimSpaces(src []rune, cursor int) int {
	for cursor < len(src) {
		if !unicode.IsSpace(src[cursor]) {
			break
		}

		cursor++
	}

	return cursor
}

func lexSyntax(src []rune, cursor int) (int, *token) {
	if src[cursor] != '(' && src[cursor] != ')' {
		return cursor, nil
	}

	return cursor + 1, &token{
		value:    string(src[cursor]),
		kind:     syntax,
		location: cursor,
	}
}

func lexInteger(src []rune, cursor int) (int, *token) {
	oldCursor := cursor

	var value []rune
	for cursor < len(src) {
		r := src[cursor]
		if r < '0' || r > '9' {
			break
		}

		value = append(value, r)
		cursor++
	}

	if len(value) == 0 {
		return oldCursor, nil
	}

	return cursor, &token{
		value:    string(value),
		kind:     integer,
		location: oldCursor,
	}
}

func lexIdentifier(src []rune, cursor int) (int, *token) {
	oldCursor := cursor

	var value []rune
	for cursor < len(src) {
		r := src[cursor]
		if unicode.IsSpace(r) {
			break
		}

		value = append(value, r)
		cursor++
	}

	if len(value) == 0 {
		return oldCursor, nil
	}

	return cursor, &token{
		value:    string(value),
		kind:     identifier,
		location: oldCursor,
	}
}
