package lexer

import (
	"BPL/lexer/lexer_token"
)

/*
This lexer function emits a TOKEN_EQUAL_SIGN then returns
the lexer for value.
*/
func LexEqualSign(lexer *Lexer) LexFn {
	lexer.Pos += len(lexertoken.ASSIGN)
	lexer.Emit(lexertoken.TOKEN_ASSIGN)
	return LexValue
}