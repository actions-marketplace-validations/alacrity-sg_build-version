package lib

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func CheckIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}

func GetIncrementType(incrementType string, offlineMode bool) (*string, error) {
	defaultIncrement := "patch"

	if incrementType != "" {
		// Check increment type.
		lowercaseIncrementType := strings.ToLower(incrementType)
		if lowercaseIncrementType == "major" || lowercaseIncrementType == "minor" || lowercaseIncrementType == "patch" {
			return &lowercaseIncrementType, nil
		} else {
			return nil, errors.New(fmt.Sprintf("Expected IncrementType to be 'major', 'minor' or 'patch' but received '%s'", lowercaseIncrementType))
		}
	}
	return &defaultIncrement, nil
}
