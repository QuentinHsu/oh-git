package utils

import (
	"fmt"

	"github.com/fatih/color"
)

// PrintWithBackgroundColor function
func PrintWithBackgroundColor(format string, backgroundColor color.Attribute, args ...interface{}) {
	color.Set(backgroundColor)
	fmt.Printf(format, args...)
	color.Unset()
}

// PrintWithColor function
func PrintWithColor(format string, textColor color.Attribute, args ...interface{}) {
	color.Set(textColor)
	fmt.Printf(format, args...)
	color.Unset()
}
