// Package cli ...
package cli

import (
	"os"
	"strings"
)

// CopyFile ...
func CopyFile(src, dst string) error {
	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = out.Write(data)
	return err
}

// RemoveLastSegment ...
func RemoveLastSegment(value string) string {
	if idx := strings.LastIndex(value, "/"); idx != -1 {
		return value[:idx]
	}
	return value
}
