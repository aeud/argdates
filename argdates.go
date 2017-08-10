package argdates

import "time"
import "fmt"

func isStringNull(s *string) bool {
	return s == nil || *s == ""
}

func isUintNull(u *uint) bool {
	return u == nil || *u == 0
}

func Days(start, end *string, duration, offset *uint) (days []time.Time, err error) {
	var d1 time.Time
	var d2 time.Time

	if !isStringNull(end) && !isStringNull(start) {
		d1, err = time.Parse("2006-01-02", *start)
		d2, err = time.Parse("2006-01-02", *end)
	} else if !isStringNull(start) && !isUintNull(duration) {
		d1, err = time.Parse("2006-01-02", *start)
		d2 = d1.AddDate(0, 0, int(*duration)-1)
	} else if !isStringNull(end) && !isUintNull(duration) {
		d2, err = time.Parse("2006-01-02", *end)
		d1 = d2.AddDate(0, 0, -int(*duration)+1)
	} else if !isUintNull(duration) {
		now := time.Now().UTC()
		d2 = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
		d1 = d2.AddDate(0, 0, -int(*duration)+1)
	} else {
		err = fmt.Errorf("Not enough parameter")
	}
	if !isUintNull(offset) {
		d1 = d1.AddDate(0, 0, -int(*offset))
		d2 = d2.AddDate(0, 0, -int(*offset))
	}
	for d := d1; d2.Sub(d).Seconds() >= 0; d = d.AddDate(0, 0, 1) {
		days = append(days, d)
	}
	return
}
