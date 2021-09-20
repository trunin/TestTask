package logger

import "fmt"

type Logger struct {

}

func NewLogger() Logger{
	return Logger{}
}

func (l *Logger)Info(text string){
	fmt.Printf("[INFO]: %s\n", text)

}

func (l *Logger)Warn(text string){
	fmt.Printf("[WARN]: %s\n", text)

}

func (l *Logger)Error(text string){
	fmt.Printf("[ERROR]: %s\n", text)
}
