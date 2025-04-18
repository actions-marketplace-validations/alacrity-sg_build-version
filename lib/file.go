package lib

import (
	"fmt"
	"os"
)

func WriteToFile(version string, outputFilePath string) error {
	data := []byte(fmt.Sprintf("BUILD_VERSION=%s", version))
	err := os.WriteFile(outputFilePath, data, 0666)
	if err != nil {
		return err
	}
}
