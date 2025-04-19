package file

import (
	"fmt"
	"os"
)

func WriteToFile(version string, outputFilePath string) error {
	data := fmt.Appendf([]byte{}, "BUILD_VERSION=%s", version)
	err := os.WriteFile(outputFilePath, data, 0666)
	if err != nil {
		return err
	}
	return nil
}
