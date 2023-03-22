package vt

import (
	"bufio"
	"os"
)

// IsStdin returns a bool determining whether standard input is not empty.
func IsStdin() bool {
	stat, _ := os.Stdin.Stat()
	return (stat.Mode() & os.ModeCharDevice) == 0
}

// ReadStdin returns a slice of strings containing text read from standard input
// using bufio scanner.
func ReadStdin() []string {
	var lines []string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
