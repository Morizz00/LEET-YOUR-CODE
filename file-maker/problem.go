package problem

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func CreateProblemFile(name, content string) error {
	safeName := strings.ReplaceAll(strings.ToLower(name), " ", "_")
	fileName := fmt.Sprintf("%s.go", safeName)
	path := filepath.Join(".", safeName)

	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}

	fullPath := filepath.Join(path, fileName)
	return os.WriteFile(fullPath, []byte(content), 0644)
}
