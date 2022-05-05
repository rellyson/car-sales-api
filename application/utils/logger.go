package utils

import (
	"os"

	"github.com/fatih/color"
)

type logger struct {
	color *color.Color
}

type Logger interface {
	Log(m string)
	Info(m string)
	Warning(m string)
	Error(m string)
	Fatal(m ...any)
	Custom(m string, c color.Attribute)
}

func NewLogger() Logger {
	return &logger{color: color.New(color.FgWhite)}
}

// Print given message using white text color (default).
func (l *logger) Log(m string) {
	l.color.Println(m)
}

// Print given message using green text color.
func (l *logger) Info(m string) {
	l.color.Add(color.FgGreen).Println(m)
}

// Print given message using yellow text color.
func (l *logger) Warning(m string) {
	l.color.Add(color.FgYellow).Println(m)
}

// Print given message using red text color.
func (l *logger) Error(m string) {
	l.color.Add(color.FgRed).Println(m)
}

// Print given message using red text color followed by an os.Exit(1) call.
func (l *logger) Fatal(m ...any) {
	l.color.Add(color.FgRed).Println(m)
	os.Exit(1)
}

func (l *logger) Custom(m string, c color.Attribute) {
	l.color.Add(c).Println(m)
}
