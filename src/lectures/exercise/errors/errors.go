//--Summary:
//  Create a function that can parse time strings into component values.
//
//--Requirements:
//* The function must parse a string into a struct containing:
//  - Hour, minute, and second integer components
//* If parsing fails, then a descriptive error must be returned
//* Write some unit tests to check your work
//  - Run tests with `go test ./exercise/errors`
//
//--Notes:
//* Example time string: 14:07:33
//* Use the `strings` package from stdlib to get time components
//* Use the `strconv` package from stdlib to convert strings to ints
//* Use the `errors` package to generate errors

// package timeparse
// TODO how to use package timeparse? get error its not a main package

package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TimeParse struct {
	Hours, Minutes, Seconds int
}

type TimeParseError struct {
	msg   string
	input string
}

func (t *TimeParseError) Error() string {
	return fmt.Sprintf("Error parsing time: %v, got %v", t.msg, t.input)
}

//attempts to pase single element to component and checks against limit (hours : 24, minutes and seconds 60)
func ParseStringToIntegerComponent(s string, limit int) (component int, err error) {

	component, err = strconv.Atoi(s)
	if err == nil && component <= limit && component >= 0 {
		return component, nil
	} else {
		return 0, &TimeParseError{fmt.Sprintf("Error converting value make sure it is number between 0 and %v", limit), s}
	}
}

func ParseDigitToTimeParse(s string) (TimeParse, error) {
	var T TimeParse
	args := strings.Split(s, ":")
	if len(args) < 3 || len(args) > 3 {
		return TimeParse{}, &TimeParseError{"Invalid number of arguments, must supply string hh:mm:ss ", s}
	} else {
		var (
			h, m, s int
		)
		var err error
		h, err = ParseStringToIntegerComponent(args[0], 24)
		if err == nil {
			T.Hours = h
		} else {
			return TimeParse{}, err
		}

		m, err = ParseStringToIntegerComponent(args[1], 60)
		if err == nil {
			T.Minutes = m
		} else {
			return TimeParse{}, err
		}

		s, err = ParseStringToIntegerComponent(args[2], 60)
		if err == nil {
			T.Seconds = s
		} else {
			return TimeParse{}, err
		}
	}
	return T, nil
}

func main() {
	fmt.Println(ParseDigitToTimeParse("112:31:11"))
}
