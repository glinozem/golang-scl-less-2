package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
)

type Logger struct {
	prefix string
}

func New(prefix string) Logger {
	return Logger{prefix: prefix}
}

func (l Logger) Info(format string, args ...any) {
	ts := time.Now().Format(time.RFC3339)
	msg := fmt.Sprintf(format, args...)
	color.New(color.FgGreen).Fprintf(os.Stdout, "%s [%s] INFO  %s\n", ts, l.prefix, msg)
}

func (l Logger) Error(format string, args ...any) {
	ts := time.Now().Format(time.RFC3339)
	msg := fmt.Sprintf(format, args...)
	color.New(color.FgRed).Fprintf(os.Stderr, "%s [%s] ERROR %s\n", ts, l.prefix, msg)
}
