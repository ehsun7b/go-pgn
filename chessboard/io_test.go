package chessboard

import (
	"testing"
	"time"
)

func TestIsTagPairGoodLines(t *testing.T) {
	lines := []string{
		"[Event \"Game\"]",
		"[Date \"2018.10.6\"]",
		"[White \"Ehsun\"]",
		"[Black \"CPU\"]",
	}

	for _, line := range lines {
		if !isTagPair(line) {
			t.Errorf("failed with value: %v", line)
		}
	}
}

func TestIsTagPairBadLines(t *testing.T) {
	lines := []string{
		"[Event Game]",
		"[Date\"2018.10.6\"]",
		"[White Ehsun]",
		"Black \"CPU\"]",
	}

	for _, line := range lines {
		if isTagPair(line) {
			t.Errorf("failed with value: %v", line)
		}
	}
}

func TestParsePairGoodLines(t *testing.T) {
	lines := []string{
		"[Event \"Game\"]",
		"[Date \"2018.10.6\"]",
		"[White \"Ehsun\"]",
		"[Black \"CPU\"]",
	}

	for _, line := range lines {
		key, value := parseTagPair(line)

		if len(key) <= 0 || len(value) <= 0 {
			t.Errorf("failed with key/value: %v/%v", key, value)
		} else {
			//log.Printf("key: %v, value: %v", key, value)
		}
	}
}

func TestParseDate1(t *testing.T) {
	value := "2018.10.6"

	d, e := parseDate(value)

	if e != nil {
		t.Errorf("Parse date for value: %v failed. Error: %v", value, e.Error())
	} else {
		expected := time.Date(2018, 10, 6, 0, 0, 0, 0, time.UTC)
		if !expected.Equal(d) {
			t.Errorf("Expected %v, parsed %v", expected, d)
		}
	}
}

func TestParseDate2(t *testing.T) {
	value := "2018.10.16"

	d, e := parseDate(value)

	if e != nil {
		t.Errorf("Parse date for value: %v failed. Error: %v", value, e.Error())
	} else {
		expected := time.Date(2018, 10, 16, 0, 0, 0, 0, time.UTC)
		if !expected.Equal(d) {
			t.Errorf("Expected %v, parsed %v", expected, d)
		}
	}
}

func TestParseDate3(t *testing.T) {
	value := "2018.2.26"

	d, e := parseDate(value)

	if e != nil {
		t.Errorf("Parse date for value: %v failed. Error: %v", value, e.Error())
	} else {
		expected := time.Date(2018, 2, 26, 0, 0, 0, 0, time.UTC)
		if !expected.Equal(d) {
			t.Errorf("Expected %v, parsed %v", expected, d)
		}
	}
}

func TestParseDate(t *testing.T) {
	value := "18.10.6"

	d, e := parseDate(value)

	if e != nil {
		t.Errorf("Parse date for value: %v failed. Error: %v", value, e.Error())
	} else {
		expected := time.Date(2018, 10, 6, 0, 0, 0, 0, time.UTC)
		if !expected.Equal(d) {
			t.Errorf("Expected %v, parsed %v", expected, d)
		}
	}
}
