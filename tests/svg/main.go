//go:build gofuzz
// +build gofuzz

package fuzz

import "github.com/jilieryuyi/canvas"

// Fuzz is a fuzz test.
func Fuzz(data []byte) int {
	_, _ = canvas.ParseSVGPath(string(data))
	return 1
}
