package date

import (
	"fmt"
	"time"
)

var months = [...]string{
	"Januari",
	"Februari",
	"Maret",
	"April",
	"Mei",
	"Juni",
	"Juli",
	"Agustus",
	"September",
	"Oktober",
	"November",
	"Desember",
}

var days = [...]string{
	"Minggu",
	"Senin",
	"Selasa",
	"Rabu",
	"Kamis",
	"Jumat",
	"Sabtu",
}

// func DateToday() *time.Time {
// 	now := time.Now()
// 	return &now
// }

// func DateTodayNow() *string {
// 	now := time.Now().Format("2006/01/02")
// 	return &now
// }

// func DateTodayEnd(month int) *string {
// 	now := time.Now().UTC().Add(time.Hour * 7)
// 	after := now.AddDate(0, month, 0).Format("2006/01/02")
// 	return &after
// }

func DateTodayLocal() *time.Time {
	now := time.Now().UTC().Add(time.Hour * 7)
	return &now
}

func DateTodayLocalWithFormat(format string) *string {
	getTime := time.Now().UTC().Add(time.Hour * 7)
	now := getTime.Format(format)
	return &now
}

func DateTodayRange() (*time.Time, *time.Time) {
	now := DateTodayLocal()
	dateStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	return &dateStart, now
}

func DateBackwardMonthRange(month int) (*time.Time, *time.Time) {
	now := DateTodayLocal()
	dateBackward := now.AddDate(0, -month, 0)
	return now, &dateBackward
}

func TimeDue() string {
	now := time.Now()
	after := now.Add(2 * time.Hour).Format("15:04")
	return fmt.Sprintf("%s, %d %s %d %s",
		days[now.Weekday()], now.Day(), months[now.Month()-1], now.Year(), after)
}

func TimeNowFormatIdn() string {
	now := time.Now()
	after := now.Format("15:04")
	return fmt.Sprintf("%s, %d %s %d %s",
		days[now.Weekday()], now.Day(), months[now.Month()-1], now.Year(), after)
}

func DateNowFormatIdn() string {
	now := time.Now()
	return fmt.Sprintf("%s, %d %s %d ",
		days[now.Weekday()], now.Day(), months[now.Month()-1], now.Year())
}

func MidtransFormatIdn(t time.Time) string {
	after := t.Format("15:04:05")
	return fmt.Sprintf("%s, %d %s %d %s",
		days[t.Weekday()], t.Day(), months[t.Month()-1], t.Year(), after)
}

func DaysBetween(a, b time.Time) int {
	if a.After(b) {
		a, b = b, a
	}

	days := -a.YearDay()
	for year := a.Year(); year < b.Year(); year++ {
		days += time.Date(year, time.December, 31, 0, 0, 0, 0, time.UTC).YearDay()
	}
	days += b.YearDay()

	return days
	// return int(b.Sub(a).Hours() / 24)
}

func TimeDifference(a, b time.Time) (year, month, day, hour, min, sec int) {
	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}
	if a.After(b) {
		a, b = b, a
	}
	y1, M1, d1 := a.Date()
	y2, M2, d2 := b.Date()

	h1, m1, s1 := a.Clock()
	h2, m2, s2 := b.Clock()

	year = int(y2 - y1)
	month = int(M2 - M1)
	day = int(d2 - d1)
	hour = int(h2 - h1)
	min = int(m2 - m1)
	sec = int(s2 - s1)

	// Normalize negative values
	if sec < 0 {
		sec += 60
		min--
	}
	if min < 0 {
		min += 60
		hour--
	}
	if hour < 0 {
		hour += 24
		day--
	}
	if day < 0 {
		// days in month:
		t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
		day += 32 - t.Day()
		month--
	}
	if month < 0 {
		month += 12
		year--
	}

	return
}

func GetTatDate(created_date time.Time, tat_days int) time.Time {
	timeNow := time.Date(created_date.Year(), created_date.Month(), created_date.Day(), 23, 59, 59, 999999999, created_date.Location())
	for i := 0; i < tat_days; i++ {
		timeNow = timeNow.AddDate(0, 0, 1)
		dayName := timeNow.Weekday().String()
		if dayName == "Saturday" || dayName == "Sunday" {
			tat_days += 1
		}
	}

	return timeNow
}
