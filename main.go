package main

import (
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

const (
	TOKEN_ILLEGAL = "ILLEGAL"
	TOKEN_EOF     = "EOF"

	// Identifiers + literals
	TOKEN_IDENT  = "IDENT"
	TOKEN_INT    = "INT"
	TOKEN_STRING = "STRING"

	// Operators
	TOKEN_ASSIGN = "="
	TOKEN_PLUS   = "+"

	// Delimiters
	TOKEN_COMMA     = ","
	TOKEN_SEMICOLON = ";"
	TOKEN_LPAREN    = "("
	TOKEN_RPAREN    = ")"
	TOKEN_LBRACE    = "{"
	TOKEN_RBRACE    = "}"
	TOKEN_COLON     = ":"

	// Keywords
	TOKEN_FUNCTION = "FUNCTION"
	TOKEN_LET      = "LET"
	TOKEN_TRUE     = "TRUE"
	TOKEN_FALSE    = "FALSE"
	TOKEN_IF       = "IF"
	TOKEN_ELSE     = "ELSE"
	TOKEN_RETURN   = "RETURN"

	// Avalon Lang specific
	TOKEN_NOTE    = "NOTE"
	TOKEN_TAG     = "TAG"
	TOKEN_CONTENT = "CONTENT"
	TOKEN_LINK    = "LINK"
)

type Token struct {
	Type  string
	Value string
}

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Avalon")

	input := widget.NewMultiLineEntry()
	input.SetPlaceHolder("Enter your note here")

	output := widget.NewLabel("Lexer output will appear here")
	output.Wrapping = fyne.TextWrapWord

	processButton := widget.NewButton("Process", func() {
		lexer := NewLexer(input.Text)
		var tokens []string
		for {
			tok := lexer.NextToken()
			tokens = append(tokens, fmt.Sprintf("%s: %s", tok.Type, tok.Value))
			if tok.Type == TOKEN_EOF {
				break
			}
		}
		output.SetText(strings.Join(tokens, "\n"))
	})

	content := container.NewVBox(
		widget.NewLabel("Welcome to Avalon!"),
		input,
		processButton,
		output,
	)

	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(600, 400))
	myWindow.ShowAndRun()
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) NextToken() Token {
	var tok Token
	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = newToken(TOKEN_ASSIGN, l.ch)
	case ';':
		tok = newToken(TOKEN_SEMICOLON, l.ch)
	case '(':
		tok = newToken(TOKEN_LPAREN, l.ch)
	case ')':
		tok = newToken(TOKEN_RPAREN, l.ch)
	case ',':
		tok = newToken(TOKEN_COMMA, l.ch)
	case '+':
		tok = newToken(TOKEN_PLUS, l.ch)
	case '{':
		tok = newToken(TOKEN_LBRACE, l.ch)
	case '}':
		tok = newToken(TOKEN_RBRACE, l.ch)
	case 0:
		tok.Value = ""
		tok.Type = TOKEN_EOF
	default:
		if isLetter(l.ch) {
			tok.Value = l.readIdentifier()
			tok.Type = lookupIdent(tok.Value)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = TOKEN_INT
			tok.Value = l.readNumber()
			return tok
		} else {
			tok = newToken(TOKEN_ILLEGAL, l.ch)
		}

	}

	l.readChar()
	return tok

}

func newToken(tokenType string, ch byte) Token {
	return Token{Type: tokenType, Value: string(ch)}
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func lookupIdent(ident string) string {
	keywords := map[string]string{
		"fn":      TOKEN_FUNCTION,
		"let":     TOKEN_LET,
		"true":    TOKEN_TRUE,
		"false":   TOKEN_FALSE,
		"if":      TOKEN_IF,
		"else":    TOKEN_ELSE,
		"return":  TOKEN_RETURN,
		"note":    TOKEN_NOTE,
		"tag":     TOKEN_TAG,
		"content": TOKEN_CONTENT,
		"link":    TOKEN_LINK,
	}
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return TOKEN_IDENT
}
