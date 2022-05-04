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

func (l *logger) Log(m string) {
	l.color.Println(m)
}

func (l *logger) Info(m string) {
	l.color.Add(color.FgGreen).Println(m)
}

func (l *logger) Warning(m string) {
	l.color.Add(color.FgYellow).Println(m)
}

func (l *logger) Error(m string) {
	l.color.Add(color.FgRed).Println(m)
}

func (l *logger) Custom(m string, c color.Attribute) {
	l.color.Add(c).Println(m)
}

func (l *logger) Fatal(m ...any) {
	l.color.Add(color.FgRed).Println(m)
	os.Exit(1)
}
