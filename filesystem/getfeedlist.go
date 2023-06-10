package filesystem

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func ReadFeedFile() ([]string, error) {
	homeDir := os.Getenv("HOME")
	filePath := filepath.Join(homeDir, ".config", "newsgo", "feeds")

	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading feeds file: %s", err)
	}

	lines := strings.Split(string(content), "\n")
	return lines, nil
}
