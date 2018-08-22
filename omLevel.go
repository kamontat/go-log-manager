package om

import "github.com/fatih/color"

var (
	// PLevelNone is force-silent in PrinterLevel
	PLevelNone = &Level{OutputType: Printer, Number: -99, String: "", Alias: []string{"none"}}
	// PLevelOneLine is 1 line print only
	PLevelOneLine = &Level{OutputType: Printer, Number: -2, String: "ONE", Alias: []string{"oneline"}}
	// PLevelSilent is silent option in PrinterLevel
	PLevelSilent = &Level{OutputType: Printer, Number: -1, String: "SLT", Alias: []string{"silent"}}
	// PLevelNormal is normal way to print
	PLevelNormal = &Level{OutputType: Printer, Number: 0, String: "NML", Alias: []string{"normal"}}
	// PLevelUnneccessary is support or fully print
	PLevelUnneccessary = &Level{OutputType: Printer, Number: 1, String: "UNC", Alias: []string{"unneccessary"}}
)

var (
	// LLevelNone equals force-silent flag
	LLevelNone = &Level{OutputType: Logger, Number: -99, String: "", Alias: []string{"none"}, Color: color.New(color.Reset)}
	// LLevelError will output as error
	LLevelError = &Level{OutputType: Logger, Number: -2, String: "ERR", Alias: []string{"error"}, Color: color.New(color.FgRed, color.Underline)}
	// LLevelWarn will output as warning
	LLevelWarn = &Level{OutputType: Logger, Number: -1, String: "WRN", Alias: []string{"warn"}, Color: color.New(color.FgYellow, color.Bold)}
	// LLevelInfo will output as information, it has flag to manage more in user layout
	LLevelInfo = &Level{OutputType: Logger, Number: 0, String: "INF", Alias: []string{"info"}, Color: color.New(color.FgGreen)}
	// LLevelLog will output as log
	LLevelLog = &Level{OutputType: Logger, Number: 1, String: "LOG", Alias: []string{"log"}, Color: color.New(color.FgBlue)}
	// LLevelDebug will output as debug, it can toggle by --debug flag
	LLevelDebug = &Level{OutputType: Logger, Number: 2, String: "DBG", Alias: []string{"debug"}, Color: color.New(color.FgMagenta)}
	// LLevelVerbose will output as verbose, it can toggle by --verbose flag
	LLevelVerbose = &Level{OutputType: Logger, Number: 3, String: "VBS", Alias: []string{"verbose"}, Color: color.New(color.Reset)}
)

// ParseLevel will parse input string to Level object
func ParseLevel(level string) *Level {
	pLevel := []*Level{PLevelNone, PLevelNormal, PLevelOneLine, PLevelSilent, PLevelUnneccessary}
	for _, p := range pLevel {
		if p.CanAlias(level) {
			return p
		}
	}

	lLevel := []*Level{LLevelNone, LLevelError, LLevelWarn, LLevelLog, LLevelInfo, LLevelDebug, LLevelVerbose}
	for _, l := range lLevel {
		if l.CanAlias(level) {
			return l
		}
	}
	return LLevelNone
}

// Level is level of output
type Level struct {
	OutputType *OutputType
	Number     int8
	String     string
	Alias      []string
	Color      *color.Color
}

// GetNumber will return number value in Level
func (l *Level) GetNumber() int8 {
	return l.Number
}

// GetColor will return color object
func (l *Level) GetColor() *color.Color {
	return l.Color
}

// GetString will return colorize string argument in LoggerLevel
func (l *Level) GetString() string {
	return l.Colorize(l.String)
}

// Colorize will change input string to colorize string
func (l *Level) Colorize(message string) string {
	if l.Color != nil {
		return l.Color.Sprint(message)
	}
	return message
}

// GetAlias will return string that represent this object
func (l *Level) GetAlias() []string {
	return l.Alias
}

// CanAlias check is input can be alias
func (l *Level) CanAlias(input string) bool {
	for _, obj := range l.Alias {
		if obj == input {
			return true
		}
	}
	return false
}

// Equal is check is level are the same
func (l *Level) Equal(o *Level) bool {
	return l.GetNumber() == o.GetNumber() && l.GetString() == o.GetString() && l == o
}
