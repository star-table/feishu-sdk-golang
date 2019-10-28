package log

import "fmt"

func Error(msg ...interface{}){
	newMsg := make([]interface{}, 1)
	newMsg[0] = "Error ::\t"
	newMsg = append(newMsg, msg...)
	fmt.Println(newMsg...)
}

func Info(msg ...interface{}){
	newMsg := make([]interface{}, 1)
	newMsg[0] = "Info ::\t"
	newMsg = append(newMsg, msg...)
	fmt.Println(newMsg...)
}

func ErrorF(msg string, args ...interface{}){
	fmt.Printf("Error::\t" + msg + "\n", args)
}

func InfoF(msg string, args ...interface{}){
	fmt.Printf("Info::\t" + msg + "\n", args)
}