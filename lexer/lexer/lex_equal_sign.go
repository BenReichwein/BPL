package lexer

import (
	"BPL/lexer/lexer_token"
)

/*
This lexer function emits a TOKEN_EQUAL_SIGN then returns
the lexer for value.
*/
func LexEqualSign(lexer *Lexer) LexFn {
	lexer.Pos += len(lexertoken.EQUAL_SIGN)
	lexer.Emit(lexertoken.TOKEN_EQUAL_SIGN)
	return LexValue
}