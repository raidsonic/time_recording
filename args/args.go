package args

import (
	"errors"
	"flag"
	"fmt"
	"regexp"
	"time"
)

// ParseArgs parses and validates command-line arguments
func ParseArgs() (string, string, string, string, error) {
	timeArg := flag.String("time", "", "Work time in 24h format (e.g., 0800-1600)")
	taskArg := flag.String("task", "", "Task description")
	dayArg := flag.String("day", "", "Day of the entry (format: YYYY-MM-DD, default: today)")
	projectArg := flag.String("project", "default", "Project name (default: 'default')")

	flag.Parse()

	if *timeArg == "" || *taskArg == "" {
		return "", "", "", "", errors.New("'time' and 'task' arguments are mandatory")
	}

	matched, err := regexp.MatchString(`^\d{4}-\d{4}$`, *timeArg)
	if err != nil || !matched {
		return "", "", "", "", errors.New("'time' must be in the format 0800-1600")
	}

	if *dayArg != "" {
		matched, err := regexp.MatchString(`^\d{4}-\d{2}-\d{2}$`, *dayArg)
		if err != nil || !matched {
			return "", "", "", "", errors.New("'day' must be in the format YYYY-MM-DD")
		}
	}

	return *timeArg, *taskArg, *dayArg, *projectArg, nil
}

// FormatTime formats the time from 0800-1600 to 08:00-16:00
func FormatTime(timeArg string) string {
	return fmt.Sprintf("%s:%s-%s:%s", timeArg[:2], timeArg[2:4], timeArg[5:7], timeArg[7:])
}

// GetDay returns the current day if dayArg is empty, otherwise returns dayArg
func GetDay(dayArg string) string {
	if dayArg == "" {
		return time.Now().Format("2006-01-02")
	}
	return dayArg
}
