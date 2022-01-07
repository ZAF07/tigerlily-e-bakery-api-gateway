package logger

import (
	"fmt"
	"log"
	"os"
)

type Logger struct {
	InfoLogger * log.Logger
	WarnLogger * log.Logger
	ErrorLogger * log.Logger
}

// !! REFACTOR !! LOGGER TO RETURN METHODS. TO USE LOGGER, WE JUST CALL errorLogger() passing in the message, err and layer it happened

func NewLogger() *Logger {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("!! ERROR LOGGING!! ")
	}

	return &Logger{
		InfoLogger: log.New(file, "[INFO] ",log.Ldate|log.Ltime|log.Lshortfile),
		WarnLogger: log.New(file, "[WARN] ",log.Ldate|log.Ltime|log.Lshortfile),
		ErrorLogger: log.New(file, "[ERROR] ",log.Ldate|log.Ltime|log.Lshortfile),
	}
}

// func (l Logger) ILogger(msg string) {
// 	l.InfoLogger.Printf("%+v", msg)
// }