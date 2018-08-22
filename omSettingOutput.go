package om

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/kamontat/go-error-manager"
)

var (
	// Logger is the output for debug or show the command process
	Logger = &OutputType{name: "Logger"}
	// Printer is the output for show the result to user/customer
	Printer = &OutputType{name: "Printer"}
)

// OutputType contain 2 variable Logger and Printer
type OutputType struct {
	name string
}

// Equal is equal method to check 2 outputType
func (o *OutputType) Equal(t OutputType) bool {
	return o.name == t.name
}

// OutputTo have 2 type; stdout and file
type OutputTo struct {
	Stdout   bool
	File     bool
	FileName string
}

// Output show output to specify location and type
func (o *OutputTo) Output(output *Output, level *Level, title string, message interface{}) {
	// Update color setting
	output.UpdateColor()

	if o.Stdout && output.IsSupport(level) {
		// Print to stdout
		for _, l := range ConstrantStandartOutputLevelList {
			if l.Equal(level) {
				fmt.Fprint(ConstrantStandartOutput, output.SPrint(level, title, message))
			}
		}

		// Print to stderr
		for _, l := range ConstrantStandartErrorLevelList {
			if l.Equal(level) {
				fmt.Fprint(ConstrantStandartError, output.SPrint(level, title, message))
			}
		}
	}

	if o.File {
		previousColor := output.GetSetting().Color
		output.DisableColor()

		exception := manager.StartErrorManager()

		err := os.MkdirAll(path.Dir(o.FileName), os.ModePerm)
		exception.AddNewError(err)
		f, err := o.GetFile()
		exception.AddNewError(err)

		if !exception.HaveError() {

			f.WriteString(output.setting.Format(output, level, title, message))
		} else {
			exception.Throw().ShowMessage()
		}

		if previousColor {
			output.EnableColor()
		}
	}

	// Reupdate color setting
	output.UpdateColor()
}

// GetFile will return File object if error ot occur and o.File set to be true
func (o *OutputTo) GetFile() (*os.File, error) {
	if o.File == false || o.FileName == "" {
		return nil, errors.New("output not support file")
	}
	p, err := filepath.Abs(o.FileName)
	if err != nil {
		return nil, err
	}
	f, err := os.OpenFile(p, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}
	return f, nil
}
