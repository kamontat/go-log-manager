package om

import (
	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
)

// FormatMessage is function to format message
type FormatMessage func(*Output, Level, string, string) string

// Setting is setting object for Output
type Setting struct {
	Type  *OutputType
	Level Level
	To    *OutputTo
	Color bool
}

// SetMaximumLevel will update maximum level of this setting
func (s *Setting) SetMaximumLevel(level Level) *Setting {
	s.Level = level
	return s
}

// SetOutputTo will update output of this setting
func (s *Setting) SetOutputTo(to *OutputTo) *Setting {
	s.To = to
	return s
}

// SetColor will update is color of this setting
func (s *Setting) SetColor(color bool) *Setting {
	s.Color = color
	return s
}

// Output setting
type Output struct {
	Name    string
	format  FormatMessage
	setting *Setting
}

// New is create new output object
func New(name string, format FormatMessage, setting *Setting) *Output {
	if setting.Type == nil {
		setting.Type = Logger
	}
	if setting.To == nil {
		setting.To = &OutputTo{Stdout: true, File: true, FileName: "/tmp/output/out-" + name + ".log"}
	}
	if format == nil {
		if setting.Type == Logger {
			format = DefaultLoggerFormatMessage
		} else if setting.Type == Printer {
			format = DefaultPrinterFormatMessage
		}
	}

	return &Output{
		Name:    name,
		format:  format,
		setting: setting,
	}
}

// DefaultLoggerFormatMessage is FormatMessage type which output message as logger formating
func DefaultLoggerFormatMessage(o *Output, level Level, title string, message string) string {
	return fmt.Sprintf("%-23s [%3s] [%-15s]: %s\n", level.Colorize(time.Now().Format(DateFormat)), level.GetString(), strings.Title(title), message)
}

// DefaultPrinterFormatMessage is FormatMessage type which output message as printer formating
func DefaultPrinterFormatMessage(o *Output, level Level, title string, message string) string {
	return fmt.Sprintf("%-15s: %s\n", strings.Title(title), message)
}

// IsSupport will check current setting is support input level
func (o *Output) IsSupport(level Level) bool {
	return o.GetSetting().Level.GetNumber() >= level.GetNumber()
}

// SPrint will return string formated
func (o *Output) SPrint(level Level, title string, message string) string {
	// Update color setting
	color.NoColor = !o.GetSetting().Color
	if o.IsSupport(level) {
		return o.format(o, level, title, message)
	}
	return ""
}

// Print will print input as string formated
func (o *Output) Print(level Level, title string, message string) {
	if o.IsSupport(level) {
		o.GetSetting().To.Output(o, level, title, message)
	}
}

// GetSetting will return setting
func (o *Output) GetSetting() Setting {
	return *o.setting
}

// DisableColor will mark output as no color
func (o *Output) DisableColor() *Output {
	o.setting.Color = false
	return o
}

// EnableColor will mark output as colorize
func (o *Output) EnableColor() *Output {
	o.setting.Color = true
	return o
}
