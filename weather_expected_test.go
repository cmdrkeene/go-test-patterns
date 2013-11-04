package main

import "testing"

func TestCurrentWeather(t *testing.T) {
	expected := "72°F, Sunny"
	actual, err := CurrentWeather("New York, NY")
	if err != nil {
		t.Error(err)
	}

	if actual != expected { // HL
		t.Error("expected", expected) // HL
		t.Error("actual  ", actual)   // HL
	} // HL
}
