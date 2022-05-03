package utils

import "github.com/fatih/color"

func NewColorizedLogger(c color.Attribute) *color.Color {
	return color.New(c)
}
