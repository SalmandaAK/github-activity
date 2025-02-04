package view

import (
	"fmt"
	"strings"
	"time"
)

const (
	oneSecond = time.Second
	oneMinute = time.Minute
	oneHour   = time.Hour
	oneDay    = 24 * time.Hour
)

func formatTimePassed(t time.Time) string {
	var msg string
	timePassed := time.Since(t)
	switch {
	case timePassed < oneSecond:
		msg = "less than a second ago"
	case timePassed < oneMinute:
		msg = "less than a minute ago"
	case timePassed < oneHour:
		msg = fmt.Sprintf("%.0f minutes ago", timePassed.Truncate(oneMinute).Minutes())
	case timePassed < oneDay:
		msg = fmt.Sprintf("%.0f hours ago", timePassed.Truncate(oneHour).Hours())
	case timePassed > oneDay:
		msg = fmt.Sprintf("%v days ago", timePassed.Truncate(oneDay).Hours()/oneDay.Hours())
	}

	// Remove 's' from minutes, hours, or days if time passed is 1
	if strings.Split(msg, " ")[0] == "1" {
		msg = strings.Replace(msg, "s", "", 1)
	}

	return msg
}
