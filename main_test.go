package main

import (
	"testing"
)

func TestSumOfValues(t *testing.T) {
	// Test with values
	values := []int{1, 2, 3}
	expectedSum := 6

	actualSum := sumOfValues(values)
	if actualSum != expectedSum {
		t.Errorf("Expected %v, got: %v", expectedSum, actualSum)
	}

	// Test when empty
	values = []int{}
	expectedSum = 0

	actualSum = sumOfValues(values)
	if actualSum != expectedSum {
		t.Errorf("Expected %v, got: %v", expectedSum, actualSum)
	}
}

func TestLowestValue(t *testing.T) {
	// Test with values
	values := []int{1, 2, 3}
	expectedSum := 1

	actualSum := lowestValue(values)
	if actualSum != expectedSum {
		t.Errorf("Expected %v, got: %v", expectedSum, actualSum)
	}

	// Test reverse order
	values = []int{3, 2, 1}
	expectedSum = 1

	actualSum = lowestValue(values)
	if actualSum != expectedSum {
		t.Errorf("Expected %v, got: %v", expectedSum, actualSum)
	}

	// Test when empty
	values = []int{}
	expectedSum = 0

	actualSum = lowestValue(values)
	if actualSum != expectedSum {
		t.Errorf("Expected %v, got: %v", expectedSum, actualSum)
	}
}

func TestDivideWithRemainder(t *testing.T) {
	var numerator, denominator int64
	var wholeNumber, remainder int64
	var expectedWholeNumber, expectedRemainder int64

	// Test with expected remainder
	numerator = 37
	denominator = 6

	expectedWholeNumber = 6
	expectedRemainder = 1

	wholeNumber, remainder = divideWithRemainder(numerator, denominator)
	if wholeNumber != expectedWholeNumber {
		t.Errorf("Expected %v, got: %v", expectedWholeNumber, wholeNumber)
	}
	if remainder != expectedRemainder {
		t.Errorf("Expected %v, got: %v", expectedRemainder, remainder)
	}

	// Test with expected no remainder
	numerator = 36
	denominator = 6

	expectedWholeNumber = 6
	expectedRemainder = 0

	wholeNumber, remainder = divideWithRemainder(numerator, denominator)
	if wholeNumber != expectedWholeNumber {
		t.Errorf("Expected %v, got: %v", expectedWholeNumber, wholeNumber)
	}
	if remainder != expectedRemainder {
		t.Errorf("Expected %v, got: %v", expectedRemainder, remainder)
	}
}

func TestGetHoursMinutesSeconds(t *testing.T) {
	var timeSecs, hours, minutes, seconds int64
	var expHours, expMinutes, expSeconds int64

	// Test with time over one hour
	timeSecs = 4213
	expHours = 1
	expMinutes = 10
	expSeconds = 13
	hours, minutes, seconds = getHoursMinutesSeconds(timeSecs)
	if hours != expHours {
		t.Errorf("Expected hours: %v, got: %v", expHours, hours)
	}
	if minutes != expMinutes {
		t.Errorf("Expected minutes: %v, got: %v", expMinutes, minutes)
	}
	if seconds != expSeconds {
		t.Errorf("Expected seconds: %v, got: %v", expSeconds, seconds)
	}

	// Test with time under one hour
	timeSecs = 613
	expHours = 0
	expMinutes = 10
	expSeconds = 13
	hours, minutes, seconds = getHoursMinutesSeconds(timeSecs)
	if hours != expHours {
		t.Errorf("Expected hours: %v, got: %v", expHours, hours)
	}
	if minutes != expMinutes {
		t.Errorf("Expected minutes: %v, got: %v", expMinutes, minutes)
	}
	if seconds != expSeconds {
		t.Errorf("Expected seconds: %v, got: %v", expSeconds, seconds)
	}

	// Test with time under one minute
	timeSecs = 15
	expHours = 0
	expMinutes = 0
	expSeconds = 15
	hours, minutes, seconds = getHoursMinutesSeconds(timeSecs)
	if hours != expHours {
		t.Errorf("Expected hours: %v, got: %v", expHours, hours)
	}
	if minutes != expMinutes {
		t.Errorf("Expected minutes: %v, got: %v", expMinutes, minutes)
	}
	if seconds != expSeconds {
		t.Errorf("Expected seconds: %v, got: %v", expSeconds, seconds)
	}
}

func TestTimeStringToSeconds(t *testing.T) {
	var timeString string
	var expectedSeconds, seconds int

	// With valid string
	timeString = "6:03.7"
	expectedSeconds = 364

	seconds, _ = timeStringToSeconds(timeString)
	if seconds != expectedSeconds {
		t.Errorf("Expected seconds: %v, got: %v", expectedSeconds, seconds)
	}

	// With longer valid string
	timeString = "1:6:03.7"
	expectedSeconds = 3964

	seconds, _ = timeStringToSeconds(timeString)
	if seconds != expectedSeconds {
		t.Errorf("Expected seconds: %v, got: %v", expectedSeconds, seconds)
	}

	// With invalid string
	timeString = "xxx"

	_, err := timeStringToSeconds(timeString)
	if err == nil {
		t.Error("Passes with invalid string\n")
	}
}

func TestProcessArgs(t *testing.T) {
	var args []string
	var seconds, expectedSeconds []int

	// With valid args
	args = []string{
		"xxx",
		"5:50.7",
		"5:56.3",
		"6:03.7",
		"5:35",
		"5:42.1",
		"6:07.2",
		"5:30.2",
		"5:39.6",
		"5:29.3",
	}

	expectedSeconds = []int{351, 356, 364, 335, 342, 367, 330, 340, 329}

	seconds = processArgs(args)
	if len(seconds) != len(expectedSeconds) {
		t.Errorf("Expected seconds length: %v does not match: %v", len(expectedSeconds), len(seconds))
	}
	for i, v := range seconds {
		if v != expectedSeconds[i] {
			t.Errorf("Expected seconds: %v but got: %v", expectedSeconds[i], v)
		}
	}

	// With invalid args
	args = []string{
		"xxx",
	}

	expectedSeconds = []int{}

	seconds = processArgs(args)
	if len(seconds) != len(expectedSeconds) {
		t.Errorf("Expected seconds length: %v does not match: %v", len(expectedSeconds), len(seconds))
	}
	for i, v := range seconds {
		if v != expectedSeconds[i] {
			t.Errorf("Expected seconds: %v but got: %v", expectedSeconds[i], v)
		}
	}

}
