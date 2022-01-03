package lexer

import (
	"strings"

	"BPL/lexer/errors"
	"BPL/lexer/lexer_token"
)

/*
This lexer function emits a TOKEN_KEY with the name of an
key that will be assigned a value.
*/
func LexKey(lexer *Lexer) LexFn {
	for {
		if strings.HasPrefix(lexer.InputToEnd(), lexertoken.EQUAL_SIGN) {
			lexer.Emit(lexertoken.TOKEN_KEY)
			return LexEqualSign
		}

		lexer.Inc()

		if lexer.IsEOF() {
			return lexer.Errorf(errors.LEXER_ERROR_UNEXPECTED_EOF)
		}
	}
}