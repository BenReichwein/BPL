package lexertoken

import (
	"fmt"
)

type Token struct {
	Type  TokenType
	Value string
}

// pre-defined Type
const (
	AND             = "&&"
	ASSIGN          = "="
	ASTERISK        = "*"
	ASTERISK_EQUALS = "*="
	BACKTICK        = "`"
	COLON           = ":"
	COMMA           = ","
	CONST           = "CONST"
	ELSE            = "ELSE"
	EOF rune        = 0
	FUNCTION        = "FUNCTION"
	GT              = ">"
	GT_EQUALS       = ">="
	IF              = "IF"
	LBRACE          = "{"
	LBRACKET        = "["
	LET             = "LET"
	LPAREN          = "("
	LT              = "<"
	LT_EQUALS       = "<="
	MINUS           = "-"
	MINUS_EQUALS    = "-="
	MINUS_MINUS     = "--"
	NEWLINE         = "\n"
	NULL            = "null"
	OR              = "||"
	PLUS            = "+"
	PLUS_EQUALS     = "+="
	PLUS_PLUS       = "++"
	POW             = "**"
	QUESTION        = "?"
	RBRACE          = "}"
	RBRACKET        = "]"
	RPAREN          = ")"
	SEMICOLON       = ";"
	SLASH           = "/"
	SLASH_EQUALS    = "/="
)

func (this Token) String() string {
	switch this.Type {
	case TOKEN_EOF:
		return "EOF"

	case TOKEN_ERROR:
		return this.Value
	}

	return fmt.Sprintf("%q", this.Value)
}