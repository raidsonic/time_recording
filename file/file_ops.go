package file

import (
	"fmt"
	"os"
	"strings"
)

// WriteEntry writes the entry to the file and handles month separation
func WriteEntry(filename, day, entry string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("opening file: %w", err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return fmt.Errorf("getting file info: %w", err)
	}

	if fileInfo.Size() == 0 {
		header := "| Date | Time | Project | Task |\n|------|------|---------|------|\n"
		if _, err := file.WriteString(header); err != nil {
			return fmt.Errorf("writing header: %w", err)
		}
	}

	var lastMonth string
	if fileInfo.Size() > 0 {
		data, err := os.ReadFile(filename)
		if err != nil {
			return fmt.Errorf("reading file: %w", err)
		}
		lines := string(data)
		lastLine := lines[strings.LastIndex(lines, "\n")+1:]
		if len(lastLine) > 0 {
			lastDate := strings.Split(lastLine, "|")[1]
			lastMonth = lastDate[1:8]
		}
	}

	currentMonth := day[:7]
	if lastMonth != "" && lastMonth != currentMonth {
		if _, err := file.WriteString("\n"); err != nil {
			return fmt.Errorf("writing empty line: %w", err)
		}
	}

	if _, err := file.WriteString(entry); err != nil {
		return fmt.Errorf("writing entry: %w", err)
	}

	return nil
}
