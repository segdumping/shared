package time

import (
	"fmt"
	"time"
)

var format = [...]string {
	"2006-01-02 15:04:05",
	"2006-1-2 15:04:05",
	"2006-01-02",
	time.RFC3339,
	"2006-01-02T15:04:05", // iso8601 without timezone
	time.RFC1123Z,
	time.RFC1123,
	time.RFC822Z,
	time.RFC822,
	time.RFC850,
	time.ANSIC,
	time.UnixDate,
	time.RubyDate,
	"2006-01-02 15:04:05.999999999 -0700 MST", // Time.String()
	"02 Jan 2006",
	"2006-01-02 15:04:05 -07:00",
	"2006-01-02 15:04:05 -0700",
	"2006-01-02 15:04:05Z07:00", // RFC3339 without T,
	time.Kitchen,
	time.Stamp,
	time.StampMilli,
	time.StampMicro,
	time.StampNano,
}

func ToTime(s string) (time.Time, error) {
	for _, f := range format {
		time, err := time.ParseInLocation(f, s, time.Local)
		if err == nil {
			return time, nil
		}
	}

	return time.Time{}, fmt.Errorf("unable to parse datetime: %s", s)
}
