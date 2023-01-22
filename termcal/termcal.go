package termcal

import (
	"fmt"
	"time"
)

type CalenderType string

const (
	CalenderTypeMonth CalenderType = "Month"
)

func GetDaysInMonth(t time.Time) int {
	nonLeapYear := map[int]int{
		1:  31,
		2:  28,
		3:  31,
		4:  30,
		5:  31,
		6:  30,
		7:  31,
		8:  31,
		9:  30,
		10: 31,
		11: 30,
		12: 31,
	}

	leapYear := map[int]int{
		1:  31,
		2:  29,
		3:  31,
		4:  30,
		5:  31,
		6:  30,
		7:  31,
		8:  31,
		9:  30,
		10: 31,
		11: 30,
		12: 31,
	}

	if t.Year()%4 == 0 {
		n, _ := leapYear[int(t.Month())]
		return n
	}

	n, _ := nonLeapYear[int(t.Month())]
	return n
}

// PrintCalender prints current calender to terminal
// CalenderType can be : M -> month
func PrintCalender(t CalenderType) {
	if t == "" {
		t = CalenderTypeMonth
	}

	calender := [7][7]string{}
	// get current date
	date := time.Now()
	date = date.Local()
	month := date.Month()
	year := date.Year()
	day := date.Day()

	// first day of month
	f := time.Date(year, month, day, 0, 0, 0, 0, date.Location())
	weekDay := f.Weekday()
	nDays := GetDaysInMonth(date)

	dayPrint := 1
	for j := int(weekDay); dayPrint <= nDays; j++ {
		i := (j + 1) / 7
		k := (j + 1) % 7
		if j == date.Day() {
			// current day is colored
			calender[i][k] = fmt.Sprintf("\x1b[6;30;42m%d\x1b[0m", dayPrint)
		} else {
			calender[i][k] = fmt.Sprintf("%d", dayPrint)
		}
		dayPrint += 1
	}

	fmt.Println("")
	fmt.Printf(" %d    %s %9s", year, date.Format("15:04:05"), month.String())
	fmt.Println()
	fmt.Println("  M", "  T ", " W ", " T ", " F ", " S ", " S ")

	for i := 0; i < 7; i++ {
		for j := 0; j < 7; j++ {
			fmt.Printf(" %2s ", calender[i][j])
		}
		fmt.Println("")
	}

}
