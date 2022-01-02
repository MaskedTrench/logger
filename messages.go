package logger

import (
	"os"
	"time"
)

func filed_format(fields map[string]string, format Format) ([]string, int) {
	var output []string
	if len(fields) == 0 {
		return output, -1
	}

	for k, v := range fields {
		str := "\t -- " + format.colors.field + k + " : \033[0m" + format.colors.content + v + "\033[0m"
		output = append(output, str)
	}

	return output, 0
}

func ctn_print(name string, msg string, fields []string, f Format) {
	b := "\t"
	dt := time.Now()

	n := b +
		" -- " + f.colors.name + name +
		"\033[0m[" + f.colors.date + dt.Format("01-02-2006") +
		"\033[0m] : " + f.colors.content + msg + "\033[0m"

	println(n)

	for _, v := range fields {
		println("\t" + v)
	}
}

func fl_print(name string, msg string, fi map[string]string, fl *os.File) {
	dt := time.Now()
	str := "-- " + name + "[" + dt.Format("01-02-2006") + "]" + " : " + msg
	fl.WriteString("\n" + str)
	for k, v := range fi {
		fl.WriteString(
			"\n        -- " + k + " : " + v,
		)
	}
}

func (l *Logger) Inform(msg string) int {
	if l.Level < InfoLvl && l.Level != NoneLvl {
		return -1
	}

	fields, _ := filed_format(l.Fields, l.Format.Info)

	ctn_print("Info", msg, fields, l.Format.Info)
	if l.File.File != nil && l.File.AllowParralels == true {
		fl_print("Info", msg, l.Fields, l.File.File)
	}
	return 0
}

func (l *Logger) Warn(msg string) int {
	if l.Level < WarningLvl && l.Level != NoneLvl {
		return -1
	}
	fields, _ := filed_format(l.Fields, l.Format.Warning)

	ctn_print("Warn", msg, fields, l.Format.Warning)
	if l.File.File != nil && l.File.AllowParralels == true {
		fl_print("Warn", msg, l.Fields, l.File.File)
	}
	return 0
}

func (l *Logger) Error(msg string) int {
	if l.Level < ErrorLvl && l.Level != NoneLvl {
		return -1
	}
	fields, _ := filed_format(l.Fields, l.Format.Error)

	ctn_print("Error", msg, fields, l.Format.Error)
	if l.File.File != nil && l.File.AllowParralels == true {
		fl_print("Error", msg, l.Fields, l.File.File)
	}
	return 0
}
