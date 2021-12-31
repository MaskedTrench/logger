package main

import "logger"

func main() {
	l, _ := logger.BuildLogger(950)
	l.EnableFiles(true)
}