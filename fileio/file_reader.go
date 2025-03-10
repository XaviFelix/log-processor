package fileio

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(filepath string) error {
	file, err := os.Open(filepath) // opens in read-only mode

	if err != nil {
		return fmt.Errorf("error opening filepath: %w", err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	// debug: print all lines
	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())

	}

	return fileScanner.Err()
}
