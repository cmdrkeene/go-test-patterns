func TestCurrentWeather(t *testing.T) {
	weather, err := CurrentWeather("New York, NY")
	if err != nil {
		t.Error(err)
	}

	if weather != "72°F, Sunny" {
		t.Error("current weather incorrect")
	}
}
