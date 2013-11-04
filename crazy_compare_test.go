func TestCurrentWeather(t *testing.T) {
  weather, _ := CurrentWeather("New York, NY")

  if weather.Temperature != 72 {
    t.Error("bad temp")
  }

  if weather.Conditions != "Sunny" {
    t.Error("bad conditions")
  }

  if weather.Observed != time.Date(2013, 11, 05, 16, 20, 0, 0, time.UTC) {
    t.Error("bad observation time")
  }
}
