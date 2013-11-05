func TestServer(t *testing.T) {
  ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, client")
  }))
  defer ts.Close()

  res, err := http.Get(ts.URL)
  if err != nil {
    log.Fatal(err)
  }
  greeting, err := ioutil.ReadAll(res.Body)
  res.Body.Close()
  if err != nil {
    log.Fatal(err)
  }

  fmt.Printf("%s", greeting)
}
