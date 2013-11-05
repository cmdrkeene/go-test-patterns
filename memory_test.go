// Adapted from src/pkg/fmt/fmt_test.go
var mallocBuf bytes.Buffer
var mallocTest = []struct {
  count int
  desc  string
  fn    func()
}{
  {0, `Sprintf("")`, func() { Sprintf("") }},
  {1, `Fprintf(buf, "%s")`, func() { mallocBuf.Reset(); Fprintf(&mallocBuf, "%s", "hello") }},
}

func TestCountMallocs(t *testing.T) {
  for _, mt := range mallocTest {
    mallocs := testing.AllocsPerRun(100, mt.fn) // HL
    if got, max := mallocs, float64(mt.count); got > max {
      t.Errorf("%s: got %v allocs, want <=%v", mt.desc, got, max)
    }
  }
}
