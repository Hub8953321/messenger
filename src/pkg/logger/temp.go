package logger

import "fmt"

type TempLogger struct {
}

var _ Logger = (*TempLogger)(nil)

func NewTempLogger() Logger {
	return &TempLogger{}
}

func (l *TempLogger) Info(msg string) {
	fmt.Println(msg)
}

func (l *TempLogger) Error(msg string) {
	fmt.Println(msg)
}

func (l *TempLogger) Fatal(msg string) {
	fmt.Println(msg)
}
