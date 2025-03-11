package parsehandling

import (
	"encoding/json"
	"fmt"
	"os"
)

type LogEntry struct {
	Level   string `json:"level"`
	Message string `json:"message"`
}

// TODO: fix this so it creates a std json file
func SaveToJSON(filepath string, logEntry LogEntry) error {
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("something went wrong opening file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(logEntry)
}
