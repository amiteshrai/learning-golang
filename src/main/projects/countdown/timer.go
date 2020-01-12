/** Build a countdown timer application */

/**
* To run the application, use:
* -> go run src/main/projects/countdown/timer.go -deadline=2020-05-03T17:35:00+01:00
 */
package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

type countdown struct {
	t int
	d int
	h int
	m int
	s int
}

func main() {
	// Declare a flag with name, default value and its description
	deadline := flag.String("deadline", "", "The deadline for the countdown timer in RFC3339 format (e.g. 2019-12-25T15:00:00+01:00)")
	flag.Parse()

	if *deadline == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	value, err := time.Parse(time.RFC3339, *deadline)
	if err != nil {
		fmt.Println("Error -> ", err)
		os.Exit(1)
	}

	fmt.Println("Remaining time to completion:")
	for range time.Tick(1 * time.Second) {
		timeRemaining := getRemainingTime(value)

		if timeRemaining.t <= 0 {
			fmt.Println("Time's up!")
			break
		}

		fmt.Printf("Days: %02d, Hours: %02d, Minutes: %02d, & Seconds: %02d ", timeRemaining.d, timeRemaining.h, timeRemaining.m, timeRemaining.s)
		fmt.Println()
	}

}

// A function to calculate the remaining time left
func getRemainingTime(t time.Time) countdown {
	currentTime := time.Now()
	diff := t.Sub(currentTime)

	total := int(diff.Seconds())
	days := int(total / (60 * 60 * 24))
	hours := int(total / (60 * 60) % 24)
	minutes := int(total/60) % 60
	seconds := int(total % 60)

	return countdown{
		t: total,
		d: days,
		h: hours,
		m: minutes,
		s: seconds,
	}
}
