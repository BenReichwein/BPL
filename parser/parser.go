package parser

import (
	"log"
	"strings"

	"BPL/model/bpl"
	"BPL/lexer"
	"BPL/lexer/lexer_token"
)

func isEOF(token lexertoken.Token) bool {
	return token.Type == lexertoken.TOKEN_EOF
}

func Parse(fileName, input string) bpl.BplFile {
	output := bpl.BplFile{
		FileName: fileName,
		Sections: make([]bpl.BplSection, 0),
	}

	var token lexertoken.Token
	var tokenValue string

	/* State variables */
	section := bpl.BplSection{}
	key := ""

	log.Println("Starting lexer and parser for file", fileName, "...")

	l := lexer.BeginLexing(fileName, input)

	for {
		token = l.NextToken()

		if token.Type != lexertoken.TOKEN_VALUE {
			tokenValue = strings.TrimSpace(token.Value)
		} else {
			tokenValue = token.Value
		}

		if isEOF(token) {
			output.Sections = append(output.Sections, section)
			break
		}

		switch token.Type {
		case lexertoken.TOKEN_SECTION:
			/*
			 * Reset tracking variables
			 */
			if len(section.KeyValuePairs) > 0 {
				output.Sections = append(output.Sections, section)
			}

			key = ""

			section.Name = tokenValue
			section.KeyValuePairs = make([]bpl.BplKeyValue, 0)

		case lexertoken.TOKEN_KEY:
			key = tokenValue

		case lexertoken.TOKEN_VALUE:
			section.KeyValuePairs = append(section.KeyValuePairs, bpl.BplKeyValue{Key: key, Value: tokenValue})
			key = ""
		}
	}

	log.Println("Parser has been shutdown")
	return output
}