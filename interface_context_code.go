var db *Database

func init() {
  db = Connect("user@localhost")
}

func Names(query string) ([]string, error) {
  rows, err := db.Select(query)
  //...
}
