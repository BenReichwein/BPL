package lexer

import (
	"BPL/lexer/lexer_token"
)

/*
This lexer function emits a TOKEN_SEMICOLON then returns
the lexer to end.
*/
func LexSemiColon(lexer *Lexer) LexFn {
	lexer.Pos += len(lexertoken.SEMICOLON)
	lexer.Emit(lexertoken.TOKEN_SEMICOLON)
	return LexValue
}