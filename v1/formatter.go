package logger

// TODO - Make possible several formats
// TODO - Make simple and modal formatter
// TODO - Make it possible to exclude all flags
// TODO - Make easy inteface for accessing formats

type Colors struct {
	name       string
	date       string
	field      string
	content    string
	helloworld string
}

type file struct {
	on          bool
	includetext bool
	path        string
}

type Format struct {
	special bool
	fileout file
	colors  Colors
	formula string
	fields  map[string]string
}

type Formats struct {
	Info    Format
	Error   Format
	Debug   Format
	Warning Format
}

var AllFormats Formats

// Builds Default format for each level
func init() {
	AllFormats.Info.special = false
	AllFormats.Error.special = false
	AllFormats.Debug.special = false
	AllFormats.Warning.special = false

	AllFormats.Info.fileout.on = false
	AllFormats.Error.fileout.on = false
	AllFormats.Debug.fileout.on = false
	AllFormats.Warning.fileout.on = false

	stdcolor := Colors{
		name:       "\033[1;34m",
		date:       "\033[2m",
		field:      "\033[1;33m",
		content:    "",
		helloworld: "---=\033[1;32m",
	}

	/*
		println("\t - Example format")
		println("\t -- " + stdcolor.name + "Name\033[0m ")
		println("\t -- " + stdcolor.date + "Date\033[0m ")
		println("\t -- " + stdcolor.field + "Field\033[0m ")
		println("\t -- " + stdcolor.content + "Content\033[0m ")
		println("\t -- " + stdcolor.helloworld + "helloworld\033[0m ")
		println("\t -- " + stdcolor.name + "LOGMSG\033[0m[" + stdcolor.date + "01.01.2000\033[0m]: " + stdcolor.content + "Simple message goes here")
		println("\t\t -- " + stdcolor.field + "ID\033[0m       : " + stdcolor.content + "Content goes here\033[0m")
		println("\t\t -- " + stdcolor.field + "Long name\033[0m: " + stdcolor.content + "Content goes here\033[0m")
	*/

	AllFormats.Info.colors = stdcolor
	AllFormats.Error.colors = stdcolor
	AllFormats.Debug.colors = stdcolor
	AllFormats.Warning.colors = stdcolor
}
