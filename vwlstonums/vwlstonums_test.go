package vwlstonums

import "testing"

func TestEncode(t *testing.T) {
	testTable := map[string]string{
		"pig? pig":      "p3g? p3g",
		"I love latin!": "3 l4v2 l1t3n!",
		"GOOD, banana":  "G44D, b1n1n1",
		"hi there":      "h3 th2r2",
		"aeiou":         "12345",
		"AEIOU":         "12345",
		"":              "",
		"11111":         "11111",
		"NHTG":          "NHTG",
	}

	for input, expected := range testTable {
		if output := Encode(input); output != expected {
			t.Errorf("Expected %v to be %v", output, expected)
		}
	}
}

func TestDecode(t *testing.T) {
	testTable := map[string]string{
		"p3g? p3g":      "pig? pig",
		"3 l4v2 l1t3n!": "i love latin!",
		"G44D, b1n1n1":  "GooD, banana",
		"h3 th2r2":      "hi there",
		"12345":         "aeiou",
		"":              "",
		"11111":         "aaaaa",
		"NHTG":          "NHTG",
	}

	for input, expected := range testTable {
		if output := Decode(input); output != expected {
			t.Errorf("Expected %v to be %v", output, expected)
		}
	}
}
