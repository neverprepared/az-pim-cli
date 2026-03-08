/*
Copyright © 2024 netr0m <netr0m@pm.me>
*/
package common

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsColorEnabledNoColor(t *testing.T) {
	t.Setenv("NO_COLOR", "1")
	assert.False(t, IsColorEnabled())
}

func TestIsColorEnabledDumbTerm(t *testing.T) {
	t.Setenv("NO_COLOR", "")
	t.Setenv("TERM", "dumb")
	assert.False(t, IsColorEnabled())
}

func TestColorFunctionsPassthroughWhenDisabled(t *testing.T) {
	// In CI/test environments stdout is not a TTY, so color is disabled.
	// Verify color functions return the input string unchanged.
	t.Setenv("NO_COLOR", "1")

	input := "hello"
	assert.Equal(t, input, Bold(input))
	assert.Equal(t, input, Green(input))
	assert.Equal(t, input, Yellow(input))
	assert.Equal(t, input, Red(input))
	assert.Equal(t, input, Cyan(input))
}

func TestColorFunctionsWrapWhenEnabled(t *testing.T) {
	// Force color on by temporarily pointing stdout to a char device.
	// Since we can't do that in a unit test, we test the colorize helper
	// indirectly by verifying ANSI codes are injected when the env is neutral
	// and stdout IS a char device (skip if not).
	fi, err := os.Stdout.Stat()
	if err != nil || (fi.Mode()&os.ModeCharDevice) == 0 {
		t.Skip("stdout is not a TTY; skipping color-enabled test")
	}

	os.Unsetenv("NO_COLOR")
	t.Setenv("TERM", "xterm-256color")

	input := "hello"
	assert.Contains(t, Bold(input), input)
	assert.Contains(t, Bold(input), "\033[")
}
