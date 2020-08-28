package logger

import (
	"io"

	"github.com/fatih/color"
)

type PrintFunc func(io.Writer, string, ...interface{})

var (
	Default PrintFunc = color.New(color.Reset).FprintfFunc()
	Blue    PrintFunc = color.New(color.FgBlue).FprintfFunc()
	Green   PrintFunc = color.New(color.FgGreen).FprintfFunc()
	Cyan    PrintFunc = color.New(color.FgCyan).FprintfFunc()
	Yellow  PrintFunc = color.New(color.FgYellow).FprintfFunc()
	Magenta PrintFunc = color.New(color.FgMagenta).FprintfFunc()
	Red     PrintFunc = color.New(color.FgRed).FprintfFunc()
)

// Logger is just a wrapper that prints stuff to STDOUT or STDERR,
// with optional color.
type Logger struct {
	Stdout  io.Writer
	Stderr  io.Writer
	Verbose bool
	Color   bool
}

// Outf prints stuff to STDOUT.
func (l *Logger) Outf(print PrintFunc, s string, args ...interface{}) {
	if len(args) == 0 {
		s, args = "%s", []interface{}{s}
	}
	if !l.Color {
		print = Default
	}
	print(l.Stdout, s+"\n", args...)
}

// VerboseOutf prints stuff to STDOUT if verbose mode is enabled.
func (l *Logger) VerboseOutf(print PrintFunc, s string, args ...interface{}) {
	if l.Verbose {
		l.Outf(print, s, args...)
	}
}

// Errf prints stuff to STDERR.
func (l *Logger) Errf(print PrintFunc, s string, args ...interface{}) {
	if len(args) == 0 {
		s, args = "%s", []interface{}{s}
	}
	if !l.Color {
		print = Default
	}
	print(l.Stderr, s+"\n", args...)
}

// VerboseErrf prints stuff to STDERR if verbose mode is enabled.
func (l *Logger) VerboseErrf(print PrintFunc, s string, args ...interface{}) {
	if l.Verbose {
		l.Errf(print, s, args...)
	}
}
