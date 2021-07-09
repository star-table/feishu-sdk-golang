package log

import (
	"log"
)

func Error(msg ...interface{}) {
	newMsg := make([]interface{}, 1)
	newMsg[0] = "Error ::\t"
	newMsg = append(newMsg, msg...)
	log.Println(newMsg...)
}

func Info(msg ...interface{}) {
	newMsg := make([]interface{}, 1)
	newMsg[0] = "Info ::\t"
	newMsg = append(newMsg, msg...)
	log.Println(newMsg...)
}

func ErrorF(msg string, args ...interface{}) {
	log.Printf("Error::\t"+msg+"\n", args...)
}

func InfoF(msg string, args ...interface{}) {
	log.Printf("Info::\t"+msg+"\n", args...)
}
