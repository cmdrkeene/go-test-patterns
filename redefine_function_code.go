var ErrUnknownLocation = errors.New("unknown location")

func GetWeather(location string) (*Weather, error) {
  code, err := findStationCode(location)
  if err != nil {
    return nil, err
  }
  return queryStation(code)
}

var findStationCode = func(location string) (code int, err error) { // HL
  res, err := http.Get(fmt.Sprintf("https://weather.com/?loc=%s", location)) // HL
  ...
} // HL
