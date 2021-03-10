package lexer

// Lexer - Struct
type Lexer struct {
	input 			string
	position 		int
	readPosition	int
	ch				byte
}

// New - Instatiates a new Lexer
func New(input string) *Lexer {
	l := &Lexer{ input: input }
	return l
}
