package main

import (
	"fmt"
	"time_recording/args"
	"time_recording/file"
	"time_recording/settings"
)

func main() {
	// Load settings
	config, err := settings.LoadSettings()
	if err != nil {
		fmt.Println("Error loading settings:", err)
		return
	}

	// Parse and validate arguments
	timeArg, taskArg, dayArg, projectArg, err := args.ParseArgs()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Use default project if not provided
	if projectArg == "default" {
		projectArg = config.ProjectDefault
	}

	// Format time to 08:00-16:00
	formattedTime := args.FormatTime(timeArg)

	// Use current day if not specified
	day := args.GetDay(dayArg)

	// Prepare the entry
	entry := fmt.Sprintf("| %s | %s | %s | %s |\n", day, formattedTime, projectArg, taskArg)

	// Write the entry to the file
	if err := file.WriteEntry(config.FilePath, day, entry); err != nil {
		fmt.Println("Error writing entry:", err)
		return
	}

	fmt.Println("Entry added successfully")
}
