import "reflect"

func TestCurrentWeather(t *testing.T) {
	expected := &Weather{
		Temperature: 72,
		Conditions:  "Sunny",
		Observed:    time.Date(2013, 11, 05, 16, 20, 0, 0, time.UTC),
	}
	actual, err := CurrentWeather("New York, NY")
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(actual, expected) { // HL
		t.Error("expected", expected)
		t.Error("actual  ", actual)
	}
}
