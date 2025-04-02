package main // tokenizer.go

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

// TokenType represents the type of token
type TokenType int

const (
	// Token types
	TOKEN_EOF TokenType = iota
	TOKEN_IDENT
	TOKEN_NUMBER
	TOKEN_FLOAT
	TOKEN_STRING
	TOKEN_PLUS
	TOKEN_MINUS
	TOKEN_STAR
	TOKEN_SLASH
	TOKEN_LPAREN
	TOKEN_RPAREN
	TOKEN_LBRACE
	TOKEN_RBRACE
	TOKEN_COMMA
	TOKEN_SEMICOLON
	TOKEN_COLON
	TOKEN_EQUALS
	TOKEN_EQ_EQUALS
	TOKEN_EXCLAMATION
	TOKEN_NOT_EQUALS
	TOKEN_LESS
	TOKEN_LESS_EQUALS
	TOKEN_GREATER
	TOKEN_GREATER_EQUALS
	TOKEN_ARROW
	TOKEN_COMMENT
	
	// Keywords
	TOKEN_FUNC
	TOKEN_RETURN
	TOKEN_IF
	TOKEN_ELSE
	TOKEN_WHILE
	TOKEN_TRUE
	TOKEN_FALSE
	TOKEN_VAR
	
	// Type keywords
	TOKEN_TYPE_INT
	TOKEN_TYPE_FLOAT
	TOKEN_TYPE_STRING
	TOKEN_TYPE_BOOL
	TOKEN_TYPE_VOID
)

// Token represents a lexical token
type Token struct {
	Type    TokenType
	Literal string
	Line    int
	Column  int
}

// Tokenizer creates tokens from source code
type Tokenizer struct {
	input        string
	position     int
	readPosition int
	ch           rune
	line         int
	column       int
}

// NewTokenizer creates a new tokenizer
func NewTokenizer(input string) *Tokenizer {
	t := &Tokenizer{
		input: input,
		line:  1,
	}
	t.readChar()
	return t
}

// readChar reads the next character and advances the position
func (t *Tokenizer) readChar() {
	if t.readPosition >= len(t.input) {
		t.ch = 0 // EOF
	} else {
		r, size := utf8.DecodeRuneInString(t.input[t.readPosition:])
		t.ch = r
		t.position = t.readPosition
		t.readPosition += size
		t.column++
	}
}

// peekChar returns the next character without advancing the position
func (t *Tokenizer) peekChar() rune {
	if t.readPosition >= len(t.input) {
		return 0 // EOF
	}
	r, _ := utf8.DecodeRuneInString(t.input[t.readPosition:])
	return r
}

// NextToken returns the next token
func (t *Tokenizer) NextToken() Token {
	var tok Token
	
	// Skip whitespace
	t.skipWhitespace()
	
	// Set token position
	tok.Line = t.line
	tok.Column = t.column
	
	switch t.ch {
	case '+':
		tok = newToken(TOKEN_PLUS, t.ch)
	case '-':
		if t.peekChar() == '>' {
			ch := t.ch
			t.readChar()
			literal := string(ch) + string(t.ch)
			tok = Token{Type: TOKEN_ARROW, Literal: literal, Line: tok.Line, Column: tok.Column}
		} else {
			tok = newToken(TOKEN_MINUS, t.ch)
		}
	case '*':
		tok = newToken(TOKEN_STAR, t.ch)
	case '/':
		if t.peekChar() == '/' {
			// Comment
			t.readChar() // skip the second '/'
			comment := t.readComment()
			tok = Token{Type: TOKEN_COMMENT, Literal: comment, Line: tok.Line, Column: tok.Column}
		} else {
			tok = newToken(TOKEN_SLASH, t.ch)
		}
	case '(':
		tok = newToken(TOKEN_LPAREN, t.ch)
	case ')':
		tok = newToken(TOKEN_RPAREN, t.ch)
	case '{':
		tok = newToken(TOKEN_LBRACE, t.ch)
	case '}':
		tok = newToken(TOKEN_RBRACE, t.ch)
	case ',':
		tok = newToken(TOKEN_COMMA, t.ch)
	case ';':
		tok = newToken(TOKEN_SEMICOLON, t.ch)
	case ':':
		tok = newToken(TOKEN_COLON, t.ch)
	case '=':
		if t.peekChar() == '=' {
			ch := t.ch
			t.readChar()
			literal := string(ch) + string(t.ch)
			tok = Token{Type: TOKEN_EQ_EQUALS, Literal: literal, Line: tok.Line, Column: tok.Column}
		} else {
			tok = newToken(TOKEN_EQUALS, t.ch)
		}
	case '!':
		if t.peekChar() == '=' {
			ch := t.ch
			t.readChar()
			literal := string(ch) + string(t.ch)
			tok = Token{Type: TOKEN_NOT_EQUALS, Literal: literal, Line: tok.Line, Column: tok.Column}
		} else {
			tok = newToken(TOKEN_EXCLAMATION, t.ch)
		}
	case '<':
		if t.peekChar() == '=' {
			ch := t.ch
			t.readChar()
			literal := string(ch) + string(t.ch)
			tok = Token{Type: TOKEN_LESS_EQUALS, Literal: literal, Line: tok.Line, Column: tok.Column}
		} else {
			tok = newToken(TOKEN_LESS, t.ch)
		}
	case '>':
		if t.peekChar() == '=' {
			ch := t.ch
			t.readChar()
			literal := string(ch) + string(t.ch)
			tok = Token{Type: TOKEN_GREATER_EQUALS, Literal: literal, Line: tok.Line, Column: tok.Column}
		} else {
			tok = newToken(TOKEN_GREATER, t.ch)
		}
	case '"':
		tok.Type = TOKEN_STRING
		tok.Literal = t.readString()
	case 0:
		tok.Literal = ""
		tok.Type = TOKEN_EOF
	default:
		if isLetter(t.ch) {
			tok.Literal = t.readIdentifier()
			tok.Type = lookupIdent(tok.Literal)
			return tok
		} else if isDigit(t.ch) {
			return t.readNumber()
		} else {
			tok = newToken(TOKEN_EOF, t.ch)
		}
	}
	
	t.readChar()
	return tok
}

// newToken creates a new token
func newToken(tokenType TokenType, ch rune) Token {
	return Token{Type: tokenType, Literal: string(ch)}
}

// readIdentifier reads an identifier
func (t *Tokenizer) readIdentifier() string {
	position := t.position
	for isLetter(t.ch) || isDigit(t.ch) {
		t.readChar()
	}
	return t.input[position:t.position]
}

// readNumber reads a number
func (t *Tokenizer) readNumber() Token {
	position := t.position
	isFloat := false
	
	for isDigit(t.ch) {
		t.readChar()
		// Check for decimal point
		if t.ch == '.' && !isFloat && isDigit(t.peekChar()) {
			isFloat = true
			t.readChar() // consume the '.'
		}
	}
	
	number := t.input[position:t.position]
	if isFloat {
		return Token{Type: TOKEN_FLOAT, Literal: number}
	}
	return Token{Type: TOKEN_NUMBER, Literal: number}
}

// readString reads a string
func (t *Tokenizer) readString() string {
	t.readChar() // skip opening quote
	position := t.position
	
	for t.ch != '"' && t.ch != 0 {
		t.readChar()
	}
	
	// Get string without closing quote
	return t.input[position:t.position]
}

// readComment reads a comment
func (t *Tokenizer) readComment() string {
	position := t.position
	
	for t.ch != '\n' && t.ch != 0 {
		t.readChar()
	}
	
	return strings.TrimSpace(t.input[position:t.position])
}

// skipWhitespace skips whitespace
func (t *Tokenizer) skipWhitespace() {
	for unicode.IsSpace(t.ch) {
		if t.ch == '\n' {
			t.line++
			t.column = 0
		}
		t.readChar()
	}
}

// isLetter checks if a character is a letter
func isLetter(ch rune) bool {
	return unicode.IsLetter(ch) || ch == '_'
}

// isDigit checks if a character is a digit
func isDigit(ch rune) bool {
	return unicode.IsDigit(ch)
}

// Keywords maps keywords to token types
var keywords = map[string]TokenType{
	"func":   TOKEN_FUNC,
	"return": TOKEN_RETURN,
	"if":     TOKEN_IF,
	"else":   TOKEN_ELSE,
	"while":  TOKEN_WHILE,
	"true":   TOKEN_TRUE,
	"false":  TOKEN_FALSE,
	"var":    TOKEN_VAR,
	"int":    TOKEN_TYPE_INT,
	"float":  TOKEN_TYPE_FLOAT,
	"string": TOKEN_TYPE_STRING,
	"bool":   TOKEN_TYPE_BOOL,
	"void":   TOKEN_TYPE_VOID,
}

// lookupIdent checks if an identifier is a keyword
func lookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return TOKEN_IDENT
}

// TokenTypeString returns a string representation of a token type
func TokenTypeString(tt TokenType) string {
	switch tt {
	case TOKEN_EOF:
		return "EOF"
	case TOKEN_IDENT:
		return "IDENT"
	case TOKEN_NUMBER:
		return "NUMBER"
	case TOKEN_FLOAT:
		return "FLOAT"
	case TOKEN_STRING:
		return "STRING"
	case TOKEN_PLUS:
		return "PLUS"
	case TOKEN_MINUS:
		return "MINUS"
	case TOKEN_STAR:
		return "STAR"
	case TOKEN_SLASH:
		return "SLASH"
	case TOKEN_LPAREN:
		return "LPAREN"
	case TOKEN_RPAREN:
		return "RPAREN"
	case TOKEN_LBRACE:
		return "LBRACE"
	case TOKEN_RBRACE:
		return "RBRACE"
	case TOKEN_COMMA:
		return "COMMA"
	case TOKEN_SEMICOLON:
		return "SEMICOLON"
	case TOKEN_COLON:
		return "COLON"
	case TOKEN_EQUALS:
		return "EQUALS"
	case TOKEN_EQ_EQUALS:
		return "EQ_EQUALS"
	case TOKEN_EXCLAMATION:
		return "EXCLAMATION"
	case TOKEN_NOT_EQUALS:
		return "NOT_EQUALS"
	case TOKEN_LESS:
		return "LESS"
	case TOKEN_LESS_EQUALS:
		return "LESS_EQUALS"
	case TOKEN_GREATER:
		return "GREATER"
	case TOKEN_GREATER_EQUALS:
		return "GREATER_EQUALS"
	case TOKEN_ARROW:
		return "ARROW"
	case TOKEN_COMMENT:
		return "COMMENT"
	case TOKEN_FUNC:
		return "FUNC"
	case TOKEN_RETURN:
		return "RETURN"
	case TOKEN_IF:
		return "IF"
	case TOKEN_ELSE:
		return "ELSE"
	case TOKEN_WHILE:
		return "WHILE"
	case TOKEN_TRUE:
		return "TRUE"
	case TOKEN_FALSE:
		return "FALSE"
	case TOKEN_VAR:
		return "VAR"
	case TOKEN_TYPE_INT:
		return "TYPE_INT"
	case TOKEN_TYPE_FLOAT:
		return "TYPE_FLOAT"
	case TOKEN_TYPE_STRING:
		return "TYPE_STRING"
	case TOKEN_TYPE_BOOL:
		return "TYPE_BOOL"
	case TOKEN_TYPE_VOID:
		return "TYPE_VOID"
	default:
		return fmt.Sprintf("UNKNOWN(%d)", tt)
	}
}