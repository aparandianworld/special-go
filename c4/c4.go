package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("module: c4")
	fmt.Println("version: 0.0.1")

	var result string
	meetingTimes := []struct {
		start, end string
	}{
		{"09:00", "10:30"},
		{"14:15", "16:00"},
		{"23:00", "23:30"},
	}

	for _, tc := range meetingTimes {
		duration, err := calculateMeetingDuration(tc.start, tc.end)
		if err != nil {
			result += fmt.Sprintf("Error: %v\n", err)
		} else {
			result += fmt.Sprintf("Meeting lasted %d minutes\n", duration)
		}
	}

	fmt.Println(result)

}

func calculateMeetingDuration(start, end string) (int, error) {

	st, err := time.Parse("15:04", start) // 24 hour format go does not use hh:mm format

	if err != nil {
		fmt.Println("Error parsing start time:", err)
		return 0, err
	}

	et, err := time.Parse("15:04", end)

	if err != nil {
		fmt.Println("Error parsing end time:", err)
		return 0, err
	}

	different := et.Sub(st)
	return int(different.Minutes()), nil
}
