package logger

import (
	"os"
	"strconv"
)

// TODO -- Make some sort of cfg file loader
// TODO -- Make some sort of format specifier

// Builds Logger for that perposes
func BuildLogger(lvl int) (l Logger, err error) {
	var lever bool = false
	for _, v := range AllLevels {
		if v == lvl {
			lever = true
		}
	}
	if !lever {
		return l, err
	}

	l.Fields = make(map[string]string)
	l.Level = lvl
	l.Format = AllFormats

	return l, nil
}

func (l *Logger) EnableFiles(flag bool) {
	if flag == false {
		l.File.File = nil
		l.File.AllowParralels = false
	}

	dirname := "log"
	latsfilename := "log"

	_, err := os.Stat(dirname)
	if err != nil {
		os.Mkdir(dirname, 0700)
	}
	os.Chdir(dirname)
	c := -1
	lvr := false
	for !lvr {
		c++
		fname := latsfilename + "-" + strconv.Itoa(c) + ".log"
		file, err := os.Open(fname)
		if err != nil {
			os.Create(fname)
			file, err = os.OpenFile(fname, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			l.File.File = file
			l.File.AllowParralels = true
			break
		}
	}
}
