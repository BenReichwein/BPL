package lexer

import (
	"BPL/lexer/lexer_token"
)

/*
This lexer function emits a TOKEN_RIGHT_BRACKET then returns
the lexer for a begin.
*/
func LexRightBracket(lexer *Lexer) LexFn {
	lexer.Pos += len(lexertoken.RBRACKET)
	lexer.Emit(lexertoken.TOKEN_RBRACKET)
	return LexBegin
}