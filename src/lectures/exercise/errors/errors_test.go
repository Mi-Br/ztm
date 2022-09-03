package main

import (
	"fmt"
	"testing"
)

func TestParseStringToIntegerComponent(t *testing.T) {
	hoursTestCases := make(map[string]bool)
	hoursTestCases["somestring"] = false
	hoursTestCases["-1"] = false
	hoursTestCases["25"] = false
	hoursTestCases["1"] = true
	hoursTestCases["01"] = true
	hoursTestCases["24"] = true

	minutesAndSecondsTestCases := make(map[string]bool)
	minutesAndSecondsTestCases["somestring"] = false
	minutesAndSecondsTestCases["-1"] = false
	minutesAndSecondsTestCases["60"] = true
	minutesAndSecondsTestCases["60"] = true
	minutesAndSecondsTestCases["25"] = true
	minutesAndSecondsTestCases["1"] = true
	minutesAndSecondsTestCases["01"] = true
	minutesAndSecondsTestCases["24"] = true

	for k, c := range hoursTestCases {
		_, err := ParseStringToIntegerComponent(k, 24)
		if c == true && err != nil || c == false && err == nil {
			t.Error(fmt.Printf("Incorrect hour convertion, ParseStringToIntegerComponent(%v) returned %v, expect %v", k, c, !c))
		}
	}

	for k, c := range minutesAndSecondsTestCases {
		_, err := ParseStringToIntegerComponent(k, 60)
		if c == true && err != nil || c == false && err == nil {
			t.Errorf("Incorrect minute or second convertion, ParseStringToIntegerComponent(%v) returned %v, expect %v", k, c, !c)
		}
	}
}

// check function handles incorrect len and incorrect type
