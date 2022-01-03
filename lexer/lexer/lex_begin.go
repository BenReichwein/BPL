package lexer

import (
	"strings"

	"BPL/lexer/lexer_token"
)

/*
This lexer function starts everything off. It determines if we are
beginning with a key/value assignment or a section.
*/
func LexBegin(lexer *Lexer) LexFn {
	lexer.SkipWhitespace()

	if strings.HasPrefix(lexer.InputToEnd(), lexertoken.LEFT_BRACKET) {
		return LexLeftBracket
	} else {
		return LexKey
	}
}