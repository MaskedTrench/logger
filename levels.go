package logger

var AllLevels [5]int

var NoneLvl int
var InfoLvl int
var WarningLvl int
var ErrorLvl int
var DebugLvl int

func init() {
	NoneLvl = 1000
	InfoLvl = 900
	WarningLvl = 800
	ErrorLvl = 700
	DebugLvl = 950

	AllLevels[0] = NoneLvl
	AllLevels[1] = InfoLvl
	AllLevels[2] = WarningLvl
	AllLevels[3] = ErrorLvl
	AllLevels[4] = DebugLvl
}

// TODO -- Make something for easy and soft changes between levels
// TODO -- Do smth for int instacing of levels, not running `log.Info`
