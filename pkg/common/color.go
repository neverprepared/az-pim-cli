/*
Copyright © 2024 netr0m <netr0m@pm.me>
*/
package common

import "os"

// IsColorEnabled returns true when ANSI color should be written to stdout.
// Respects https://no-color.org and detects non-terminal output.
func IsColorEnabled() bool {
	if os.Getenv("NO_COLOR") != "" || os.Getenv("TERM") == "dumb" {
		return false
	}
	fi, err := os.Stdout.Stat()
	if err != nil {
		return false
	}
	return (fi.Mode() & os.ModeCharDevice) != 0
}

func colorize(s, code string) string {
	if !IsColorEnabled() {
		return s
	}
	return "\033[" + code + "m" + s + "\033[0m"
}

func Bold(s string) string   { return colorize(s, "1") }
func Green(s string) string  { return colorize(s, "32") }
func Yellow(s string) string { return colorize(s, "33") }
func Red(s string) string    { return colorize(s, "31") }
func Cyan(s string) string   { return colorize(s, "36") }
