package shared

type Token string

func nullToken() Token { return "" }

func isNullToken(t Token) bool {
	return t.Kind() == 0
}

func (tk Token) Kind() OpKind {
	if len(tk) == 0 {
		return 0
	}

	// since tokens in BF are groups of at least
	// one character of the same type. we can take
	// the first char as a representative
	return OpKind(tk[0])
}

type Tokenizer struct{}

func (t *Tokenizer) Tokenize(content []byte) []Token {
	var tokens []Token

	currToken := nullToken()

	for _, c := range content {
		if !isNullToken(currToken) && currToken.Kind() != OpKind(c) {
			tokens = append(tokens, currToken)
			currToken = nullToken()
		}

		switch OpKind(c) {
		case OpPlus, OpMinus, OpLeftArrow, OpRightArrow:
			currToken += Token(c)
		case OpDot, OpLeftBracket, OpRightBracket:
			currToken += Token(c)
			tokens = append(tokens, currToken)
			currToken = nullToken()
		default:
			// ignore any other characters
			continue
		}
	}

	return tokens
}
