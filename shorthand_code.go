type Token rune

const (
  Keyword Token = iota
  Wildcard
  Identifier
  EOF
)

type Lexer struct {
  source  string
}

func (l *Lexer) scan() []Token {
  tokens := []Token{}
  var t Token
  for t != EOF {
    t = l.tokenize(l.next())
    tokens = append(tokens, t)
  }
  return tokens
}
