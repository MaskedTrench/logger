package logger

type Logger struct {
	to_file      bool
	level        int
	format       string
	path_to_file string
}

func (l *Logger) isLeveled(lvl int) bool {
	if l.level <= lvl {
		return true
	}
	return false
}

// TODO - Make auto tests
// TODO - Make format embed
func (l *Logger) Info(msg string) {
	if l.isLeveled(InfoLvl) {
		println(msg)
	}
}

func (l *Logger) Warn(msg string) {
	if l.isLeveled(WarningLvl) {
		println(msg)
	}
}

func (l *Logger) Error(msg string) {
	if l.isLeveled(ErrorLvl) {
		println(msg)
	}
}

func (l *Logger) Debug(msg string) {
	if l.isLeveled(DebugLvl) {
		println(msg)
	}
}
