package fileio

import (
	"fmt"
	"os"
)

func WriteFile(filepath, contnet string) error {
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	_, err = file.WriteString(contnet + "\n")
	return err
}
