package booking

import (
	"fmt"
	"strconv"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"
	formattedT, _ := time.Parse(layout, date)
	return formattedT
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
	formattedDate, _ := time.Parse(layout, date)
	return formattedDate.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	formattedDate, _ := time.Parse(layout, date)
	hour := formattedDate.Hour()
	if hour >= 12 && hour < 18 {
		return true
	}
	return false
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	t := Schedule(date)
	formattedDate := t.Format("Monday, January 2, 2006")
	formattedTime := t.Format("15:04")
	return fmt.Sprintf("You have an appointment on %v, at %v.", formattedDate, formattedTime)
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	year := time.Now().Year()
	anniversaryDate := strconv.Itoa(year) + "-09-15"
	date, _ := time.Parse("2006-01-02", anniversaryDate)
	return date
}
