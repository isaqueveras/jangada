// Package cli ...
package cli

import (
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
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

func GetEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// Capitalize capitalizes the first letter of a string
func Capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

// GenerateTemplate generates the template files using templ
func GenerateTemplate() {
	log := color.New()
	log.Add(color.Bold, color.FgHiBlue).Print("\nGenerating template...\n\n")
	log.Add(color.FgHiGreen).Print("\trun\t")
	log.Add(color.Reset).Println(`go run github.com/a-h/templ/cmd/templ@latest generate`)

	command := exec.Command("bash", "-c", "go run github.com/a-h/templ/cmd/templ@latest generate")
	command.Stdout, command.Stderr = os.Stdout, os.Stderr
	if err := command.Run(); err != nil {
		panic(err)
	}
}
