package om

import (
	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
)

// Output setting
type Output struct {
	setting *Setting
}

// New is create new output object
func New(setting *Setting) *Output {
	setting.MarkDefault()

	return &Output{
		setting: setting,
	}
}

// DefaultLoggerFormatMessage is FormatMessage type which output message as logger formating
func DefaultLoggerFormatMessage(o *Output, level *Level, title string, message interface{}) string {
	return fmt.Sprintf("%-23s [%3s] [%-"+ConstrantLogTitleSize+"s]: %s\n", level.Colorize(time.Now().Format(ConstrantDateFormat)), level.GetString(), strings.Title(title), message)
}

// DefaultPrinterFormatMessage is FormatMessage type which output message as printer formating
func DefaultPrinterFormatMessage(o *Output, level *Level, title string, message interface{}) string {
	return fmt.Sprintf("%-"+ConstrantPrintTitleSize+"s: %s\n", strings.Title(title), message)
}

// IsSupport will check current setting is support input level
func (o *Output) IsSupport(level *Level) bool {
	return o.GetSetting().Level.GetNumber() >= level.GetNumber()
}

// UpdateColor will set color disable by output setting
func (o *Output) UpdateColor() {
	color.NoColor = !o.GetSetting().Color
}

// SPrint will return string formated
func (o *Output) SPrint(level *Level, title string, message interface{}) string {
	// Update color setting
	o.UpdateColor()

	if o.IsSupport(level) {
		return o.setting.Format(o, level, title, message)
	}
	return ""
}

// Print will print input as string formated
func (o *Output) Print(level *Level, title string, message interface{}) {
	o.GetSetting().To.Output(o, level, title, message)
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
