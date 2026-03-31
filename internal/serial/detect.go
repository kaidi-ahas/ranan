package serial

import (
	"fmt"
	"path/filepath"
)

func DetectPort(pattern string) (string, error) {
	matches, err := filepath.Glob(pattern)
	if err != nil {
		return "", fmt.Errorf("failed to scan for Arduino port: %w", err)
	}

	if len(matches) == 0 {
		return "", fmt.Errorf("no Arduino port found matching %s", pattern)
	}

	return matches[0], nil
}