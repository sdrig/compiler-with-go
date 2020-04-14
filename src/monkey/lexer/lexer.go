package lexer

import (
	"monkey/token"
)

// Lexer
type Lexer struct {
	input        string
	position     int  // current position in input(points to current char)
	readPosition int  //current reading position in input (after current char)
	ch           byte //current char under examination
}

// initialize new Lexer
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// read the current position's char
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

// bring next token
func (l *Lexer) NextToken() token.Token {
	var toke token.Token
	l.skipWhitespace()
	switch l.ch {
	case '=':
		toke = newToken(token.ASSIGN, l.ch)
	case ';':
		toke = newToken(token.SEMICOLON, l.ch)
	case '(':
		toke = newToken(token.LPAREN, l.ch)
	case ')':
		toke = newToken(token.RPAREN, l.ch)
	case ',':
		toke = newToken(token.COMMA, l.ch)
	case '+':
		toke = newToken(token.PLUS, l.ch)
	case '{':
		toke = newToken(token.LBRACE, l.ch)
	case '}':
		toke = newToken(token.RBRACE, l.ch)
	case 0:
		toke.Literal = ""
		toke.Type = token.EOF
	default:
		if isLetter(l.ch) {
			toke.Literal = l.readIdentifier()
			toke.Type = token.LookupIdent(toke.Literal)
			return toke
		} else if isDigit(l.ch) {
			toke.Type = token.INT
			toke.Literal = l.readNumber()
			return toke
		}
		toke = newToken(token.ILLEGAL, l.ch)

	}
	l.readChar()
	return toke
}
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
