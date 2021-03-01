package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func sumOfValues(values []int) int {
	sum := 0
	for _, t := range values {
		sum += t
	}

	return sum
}

func lowestValue(values []int) int {
	if len(values) < 1 {
		return 0
	}

	lowest := values[0]

	for _, number := range values {
		if number < lowest {
			lowest = number
		}
	}

	return lowest
}

func divideWithRemainder(numerator, denominator int64) (whole, remainder int64) {
	whole = numerator / denominator
	remainder = numerator % denominator
	return
}

func getHoursMinutesSeconds(timeSeconds int64) (hours, minutes, seconds int64) {
	hourSeconds := 3600
	minSeconds := 60

	hours, remainder := divideWithRemainder(timeSeconds, int64(hourSeconds))
	minutes, seconds = divideWithRemainder(remainder, int64(minSeconds))
	return
}

func validateStringToInt(value string) int {
	fVal, err := strconv.ParseFloat(value, 64)
	if err != nil {
		fmt.Println("Non numeric value passed - check your splits")
		fmt.Printf("You entered: '%v'\n", value)
		os.Exit(1)
	}
	intVal := math.Round(fVal)

	return int(intVal)
}

func timeStringToSeconds(timeString string) (int, error) {

	parts := strings.Split(timeString, ":")

	var hours, mins, seconds int

	if len(parts) > 3 || len(parts) < 2 {
		return 0, errors.New("This is not a valid time")
	} else if len(parts) == 3 {
		hours, mins, seconds = validateStringToInt(parts[0]), validateStringToInt(parts[1]), validateStringToInt(parts[2])
	} else {
		hours, mins, seconds = 0, validateStringToInt(parts[0]), validateStringToInt(parts[1])
	}

	return hours*3600 + mins*60 + seconds, nil
}

func processArgs(args []string) []int {
	var allSplits []int

	for _, kilometer := range args[1:] {
		secs, err := timeStringToSeconds(kilometer)
		if err != nil {
			fmt.Println("This is not a valid time")
			fmt.Printf("You entered: '%v'\n", kilometer)
			os.Exit(1)
		}
		allSplits = append(allSplits, secs)
	}

	return allSplits
}

func main() {
	allSplits := processArgs(os.Args)

	if len(allSplits) < 5 {
		fmt.Println("You have not entered enough 1k splits to complete 5k")
		fmt.Printf("You entered: '%v'\n", os.Args[1:])
		fmt.Println("\nPlease at least five 1k splits seperated by a space: 12:42.1 4:18:26.8 ... etc")
		os.Exit(1)
	}

	var fiveK []int

	count := 1
	for i := 0; i < len(allSplits); i++ {
		if count > 4 {
			timez := allSplits[count-5 : count]
			fiveKTime := sumOfValues(timez)
			fiveK = append(fiveK, fiveKTime)
		}

		count++
	}

	fastest := lowestValue(fiveK)

	hours, minutes, seconds := getHoursMinutesSeconds(int64(fastest))

	fmt.Printf("Fastest 5k section took: %v hours, %v minutes, %v seconds\n", hours, minutes, seconds)
	fmt.Printf("%v:%v:%v\n", hours, minutes, seconds)
}
