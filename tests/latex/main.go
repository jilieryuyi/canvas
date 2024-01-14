//go:build gofuzz
// +build gofuzz

package fuzz

import "github.com/jilieryuyi/canvas"

// Fuzz is a fuzz test.
func Fuzz(data []byte) int {
	_, _ = canvas.ParseLaTeX(string(data))
	return 1
}
