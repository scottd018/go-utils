package directory

import (
	"fmt"
	"os"
)

// Exists simply returns if a directory exists and if it is actually a directory.
func Exists(path string) (bool, error) {
	// check if the directory exists
	dirInfo, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false, nil
	}

	// check if it is actually a directory
	if dirInfo.IsDir() {
		return true, nil
	}

	return false, fmt.Errorf("path [%s] exists but is not a directory", path)
}
