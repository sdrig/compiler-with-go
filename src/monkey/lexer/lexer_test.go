package lexer

import (
	"testing"
)

func TestNextToken(t *testing.T) {

	t.Run("test 1", func(t *testing.T) {
		input := `let five = 5;
	let ten = 10;
	
	let add = fn(x,y){ 
		x+y;		
	};
	let result = add(five,ten);
	`

		tests := []struct {
			expextedType    token.TokenType
			expectedLiteral string
		}{

			{token.LET, "let"},
			{token.IDENT, "five"},
			{token.ASSIGN, "="},
			{token.INT, "5"},
			{token.SEMICOLON, ";"},
			{token.LET, "let"},
			{token.IDENT, "ten"},
			{token.ASSIGN, "="},
			{token.INT, "10"},
			{token.SEMICOLON, ";"},
			{token.LET, "let"},
			{token.IDENT, "add"},
			{token.ASSIGN, "="},
			{token.FUNCTION, "fn"},
			{token.LPAREN, "("},
			{token.IDENT, "x"},
			{token.COMMA, ","},
			{token.IDENT, "y"},
			{token.RPAREN, ")"},
			{token.LBRACE, "{"},
			{token.IDENT, "x"},
			{token.PLUS, "+"},
			{token.IDENT, "y"},
			{token.SEMICOLON, ";"},
			{token.RBRACE, "}"},
			{token.SEMICOLON, ";"},
			{token.LET, "let"},
			{token.IDENT, "result"},
			{token.ASSIGN, "="},
			{token.IDENT, "add"},
			{token.LPAREN, "("},
			{token.IDENT, "five"},
			{token.COMMA, ","},
			{token.IDENT, "ten"},
			{token.RPAREN, ")"},
			{token.SEMICOLON, ";"},
			{token.EOF, ""},
		}

		l := New(input)

		for i, tt := range tests {
			tok := l.NextToken()
			// fmt.Printf("%#v\n", tok)
			if tok.Type != tt.expextedType {
				t.Fatalf("tests[%d] - tokenType wrong. expected:%q, got:%q", i, tt.expextedType, tok.Type)
			}
			if tok.Literal != tt.expectedLiteral {
				t.Fatalf("tests[%d], expected:%q, got:%q", i, tt.expectedLiteral, tok.Literal)
			}
		}
	})
	t.Run("t extended", func(t *testing.T) {
		input := `let five = 5;
let ten = 10;
let add = fn(x, y) {
x + y;
};
let result = add(five, ten);
!-/*5;
5 < 10 > 5;`
	})
}
