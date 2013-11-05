func TestUnknownLocation(t *testing.T) {
  findStationCode = func(location string) (code int, err error) { // HL
    if location == "New York, NY" { // HL
      return 10011, nil // HL
    } // HL
    return 0, ErrUnknownLocation // HL
  }

  _, err := GetWeather("Batman, Turkey")
  if err != ErrUnknownLocation {
    t.Error("we should not have found Batman")
    t.Error(err)
  }

  _, err = GetWeather("New York, NY")
  if err != nil {
    t.Error("we should have found New York, c'mon!")
    t.Error(err)
  }
}
