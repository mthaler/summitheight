package main

import (
    "testing"
)

func TestToSlice(t *testing.T) {
    s := Summit{
		Name: "Hochfelln",
		Height: "1674",
		Category: "Berg, Gipfel",
		Country: "Deutschland",
		Region: "Bayern",
		Group: "Chiemgauer Alpen",
		Information: "",
	}
	want := []string{"Hochfelln", "1674", "Berg, Gipfel","Deutschland","Bayern","Chiemgauer Alpen",""}
    if !want.MatchString(msg) || err != nil {
        t.Errorf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
    }