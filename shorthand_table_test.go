func TestScan(t *testing.T) {
	cases := []Case{
		Case{
			`SELECT * FROM users`,
			[]Token{K, W, K, I, E},
		},
		Case{
			`SELECT * FROM users LIMIT 10 ORDER BY id ASC`,
			[]Token{K, W, K, I, K, N, K, K, I, K, E},
		},
		Case{
			`SELECT id, name FROM users WHERE id = 1 AND name BETWEEN("a", "z")`,
			[]Token{K, I, C, I, K, I, K, I, O, N, O, I, O, L, S, C, S, R, E},
		},
	}
}
