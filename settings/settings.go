package settings

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Settings struct {
	ProjectDefault string `json:"project-default"`
	FilePath       string `json:"filepath"`
}

func LoadSettings() (*Settings, error) {
	execPath, err := os.Executable()
	if err != nil {
		return nil, fmt.Errorf("getting executable path: %w", err)
	}
	configPath := filepath.Join(filepath.Dir(execPath), "time_recording.config.json")

	file, err := os.Open(configPath)
	if err != nil {
		return nil, fmt.Errorf("opening config file: %w", err)
	}
	defer file.Close()

	var settings Settings
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&settings); err != nil {
		return nil, fmt.Errorf("decoding config file: %w", err)
	}

	// Expand the file path if it starts with ~
	if settings.FilePath[:2] == "~/" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return nil, fmt.Errorf("getting user home directory: %w", err)
		}
		settings.FilePath = filepath.Join(homeDir, settings.FilePath[2:])
	}

	fmt.Println(settings.FilePath)
	return &settings, nil
}
