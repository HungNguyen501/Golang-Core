package utils

import (
	"errors"
	"time"
)

// Returns timenow in specific format
func TimeNow(format string) (string, error) {
	if format == "" {
		return time.Now().Format(time.ANSIC), nil
	}
	val, ok := TimeFormats[format]
	if !ok {
		return "", errors.New("invalid time format")
	}
	return time.Now().Format(val), nil
}

// Returns the list of available time formats in system
func AvailableTimeFormats() string {
	availFormats := ""
	for k := range TimeFormats {
		availFormats += (k + " ")
	}
	return availFormats
}
