package main

import (
	"fmt"
	"time"
)

// Schedule takes in the date string and returns a time.Time
// 7/25/2019 13:45:00
func Schedule(date string) time.Time {
	t, err := time.Parse("1/02/2006 15:04:05", date)

	if err != nil {
		panic(err)
	}
	return t
}

// HasPassed function will compare the entered date to the current date and returns true or false
// July 25, 2019 13:45:00
func HasPassed(date string) bool {
	d, err := time.Parse("January 2, 2006 15:04:05", date)
	if err != nil {
		panic(err)
	}

	return d.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon
// Thursday, July 25, 2019 13:45:00
func IsAfternoonAppointment(date string) bool {
	d, err := time.Parse("Monday, January 2, 2006 15:04:05", date)
	if err != nil {
		panic(err)
	}
	if d.Hour() >= 12 && d.Hour() < 18 {
		return true
	}
	return false
}

// Description function returns a formatted string of the appointment time
// In: 7/25/2019 13:45:00
// Out: You have an appointment on Thursday, July 25, 2019, at 13:45.
func Description(date string) string {
	d, err := time.Parse("1/2/2006 15:04:05", date)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("You have an appointment on %s", d.Format("Monday, January 2, 2006, at 15:04."))
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	t, _ := time.Parse("2006-01-2", fmt.Sprintf("%d-09-15", time.Now().Year()))
	return t
}

func main() {

	fmt.Println(Schedule("9/15/2021 19:15:00"))

	fmt.Println("Here is the next value")

	fmt.Println(HasPassed("July 25, 2049 13:45:00"))

	fmt.Println("Here is the next value")

	fmt.Println(IsAfternoonAppointment("Thursday, July 25, 2019 13:45:00"))

	fmt.Println("Here is the next value")

	fmt.Println(Description("9/28/2031 09:15:00"))

	fmt.Println("Here is the next value")

	fmt.Println(AnniversaryDate())

}
