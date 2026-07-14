package logger

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// DefaultTimeLayout is layout for convert time.Now to String.
const DefaultTimeLayout = "2006-01-02 15:04:05"

// DefaultLineFormat is format to log line record.
const DefaultLineFormat = "[%s] %s: %s\n%s\n\n"

// ErrEmptyFileName is returned when trying to open file with empty name.
var ErrEmptyFileName = errors.New("empty file name")

// OpenFile opens file for append strings. Creates file and its parent
// directory if they do not exist.
func OpenFile(name string) (*os.File, error) {
	if name == "" {
		return nil, ErrEmptyFileName
	}

	if dir := filepath.Dir(name); dir != "." {
		const dirPerm = 0o755

		if err := os.MkdirAll(dir, dirPerm); err != nil {
			return nil, fmt.Errorf("create directory: %w", err)
		}
	}

	const filePerm = 0o644

	file, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, filePerm)
	if err != nil {
		return nil, fmt.Errorf("open: %w", err)
	}

	return file, nil
}

// Write saves request and response to log file.
func Write(name string, address string, request string, response string) error {
	// Disable logging if log file name is empty.
	if name == "" {
		return nil
	}

	file, err := OpenFile(name)
	if err != nil {
		return err
	}
	defer file.Close()

	line := fmt.Sprintf(DefaultLineFormat, time.Now().Format(DefaultTimeLayout), address, request, response)
	if _, err = file.WriteString(line); err != nil {
		return fmt.Errorf("write: %w", err)
	}

	return nil
}
