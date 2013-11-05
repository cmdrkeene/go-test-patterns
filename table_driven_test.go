// Adapted from src/pkg/fmt/fmt_test.go
var fmttests = []struct {
  fmt string
  val interface{}
  out string
}{
  {"%d", 12345, "12345"},
  {"%v", 12345, "12345"},
  {"%t", true, "true"},
}

func TestSprintf(t *testing.T) {
  for _, tt := range fmttests {
    s := Sprintf(tt.fmt, tt.val)
    if s != tt.out {
      t.Errorf("Sprintf(%q, %v) = %q want %q", tt.fmt, tt.val, s, tt.out)
    }
  }
}
