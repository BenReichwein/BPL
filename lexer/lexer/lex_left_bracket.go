package lexer

import (
	"BPL/lexer/lexer_token"
)

/*
This lexer function emits a TOKEN_LBRACKET then returns
the lexer for a section header.
*/
func LexLeftBracket(lexer *Lexer) LexFn {
	lexer.Pos += len(lexertoken.LBRACKET)
	lexer.Emit(lexertoken.TOKEN_LBRACKET)
	return LexSection
}