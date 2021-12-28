package logger

var AllLevels []int

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
}

// TODO -- Base values for every `level`
// TODO -- Conversetion from `Level` to `int`
// TODO -- Make something for easy and soft changes between levels
// TODO -- Do smth for int instacing of levels, not running `log.Info`
