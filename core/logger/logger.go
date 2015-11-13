package logger

import "log"

// Logger struct
type Logger struct{}

// NewLogger pointer of Logger
func NewLogger() *Logger {
	return &Logger{}
}

// Log logs to the terminal
func (Logger) Log(message string, err error) {
	if err != nil {
		log.Println(message, "Error:", err.Error())
	} else {
		log.Println(message)
	}
}
