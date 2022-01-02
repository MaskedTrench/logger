package logger

import "os"

type fs struct {
	File           *os.File
	AllowParralels bool
}

type Logger struct {
	File   fs
	Level  int
	Format Formats
	Fields map[string]string
}
