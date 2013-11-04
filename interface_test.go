type TestDB struct {
  OnSelect func(query string) ([]Row, error)
}

func (db *TestDB) Select(query string) ([]Row, error) {
  return db.OnSelect(query)
}

func TestFinderNamesError(t *testing.T) {
  expected := errors.New("statement invalid")
  db := &TestDB{
    OnSelect: func(query) ([]Row, err) {
      return nil, expected
    },
  }

  finder := &Finder{db}
  _, err := finder.Names()
  if err != expected {
    t.Error("expected", expected)
    t.Error("actual  ", err)
  }
}
