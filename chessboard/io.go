package chessboard

import (
	"errors"
	"log"
	"strings"
	"time"
)

var (
	errorEmptyString = errors.New("the string is empty")
	errorBadDate     = errors.New("bad date value")
)

func Parse(pgb string) (State, error) {
	if len(pgb) <= 0 {
		return State{}, errorEmptyString
	}

	lines := strings.Split(pgb, "\n")

	var state State

	for _, line := range lines {
		line := strings.TrimSpace(line)

		if isTagPair(line) {
			tagName, tagValue := parseTagPair(line)

			log.Println(tagName, tagValue)

			switch tagName {
			case TagEvent:
				state.event = tagValue
				break
			case TagSite:
				state.site = tagValue
				break
			case TagDate:
				date, e := parseDate(tagValue)
				state.date = date
				if e != nil {
					log.Fatalf("date parse error '%v'", e.Error())
				}
				break
			case TagRound:
				state.round = tagValue
				break
			case TagWhite:
				state.white = tagValue
				break
			case TagBlack:
				state.black = tagValue
				break
			case TagResult:
				state.result = tagValue
			}
		}
	}

	return state, nil
}

func parseDate(vale string) (time.Time, error) {
	layout, e := findDateLayout(vale)

	if e == nil {
		date, e := time.Parse(layout, vale)
		if e == nil {
			return date, nil
		} else {
			return time.Time{}, e
		}
	} else {
		return time.Time{}, e
	}
}

func findDateLayout(value string) (string, error) {
	value = strings.TrimSpace(value)

	if len(value) < 6 {
		return "", errorBadDate
	}

	parts := strings.Split(value, ".")

	if len(parts) != 3 {
		return "", errorBadDate
	}

	yearLayout := "2006"
	monthLayout := "01"
	dayLayout := "02"

	if len(parts[0]) == 2 {
		yearLayout = "06"
	}

	if len(parts[1]) == 1 {
		monthLayout = "1"
	}

	if len(parts[2]) == 1 {
		dayLayout = "2"
	}

	return yearLayout + "." + monthLayout + "." + dayLayout, nil
}

func parseTagPair(line string) (string, string) {
	line = strings.TrimSpace(strings.Trim(line, "[]"))
	parts := strings.Split(line, " ")
	return parts[0], strings.Trim(parts[len(parts)-1], "\"")
}

func isTagPair(line string) bool {
	if len(line) <= 0 {
		return false
	}

	if strings.Count(line, "\"") != 2 {
		return false
	}

	if !strings.Contains(line, " ") {
		return false
	}

	if line[0] != '[' || line[len(line)-1] != ']' {
		return false
	}

	return true
}
