package main

import (
	"TestTask/app"
	"log"
	"runtime/debug"
	"strings"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Fatalf(
				"panic occurred: %s: stack trace: %s",
				r,
				strings.Replace(string(debug.Stack()), "\n", "", -1),
			)
		}
	}()

	server := app.NewServer()
	server.Serve()
}
