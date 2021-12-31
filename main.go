package main

import "logger"

func main() {
	l, _ := logger.BuildLogger(950)
	l.EnableFiles(true)
	l.Inform("Information messages")
	l.Warn("WARNING!")
	l.Error("Error message")
}
