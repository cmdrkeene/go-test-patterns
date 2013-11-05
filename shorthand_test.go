const (
  K = Keyword
  W = Wildcard
  I = Identifier
  E = EOF
)

func TestScan(t *testing.T) {
  l := &Lexer{`SELECT * FROM users`}
  expected := []Token{K, W, K, I, E}
  actual := l.scan()
  if !reflect.DeepEqual(expected, actual) {
    t.Error("expected", expected)
    t.Error("actual  ", actual)
    t.Error("source  ", l.source)
  }
}
