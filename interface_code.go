type Finder struct {
  db *Database
}

type Database interface {
  Select(query) ([]Row, error)
}

func (f *Finder) Names() ([]string, error) {
  rows, err := f.db.Select("SELECT name FROM users")
  if err != nil {
    return nil, err
  }
  return f.process(rows), nil
}
