package om

import (
	"errors"

	manager "github.com/kamontat/go-error-manager"
)

// NewPrinter will create printer by setting
func NewPrinter(setting *Setting) *Output {
	setting.Type = Printer
	return New(setting)
}

// PrintWrapper for wrap the output to customize next chain
type PrintWrapper struct {
	out *manager.ResultWrapper
}

// Check will check output object in wrapper
func (p *PrintWrapper) Check() *manager.Throwable {
	return p.out.Catch(func() error {
		return errors.New("printer is not exist, please call SetupPrinter first")
	}, nil)
}

// unwrap will call force unwrap so be aware
func (p *PrintWrapper) unwrap() *Output {
	return p.out.ForceUnwrap().(*Output)
}

// ToOneLine print as oneline
func (p *PrintWrapper) ToOneLine(title string, message interface{}) {
	GetOutput(p).Print(PLevelOneLine, title, message)
}

// ToSilent print as slient
func (p *PrintWrapper) ToSilent(title string, message interface{}) {
	GetOutput(p).Print(PLevelSilent, title, message)
}

// ToNormal print as normal
func (p *PrintWrapper) ToNormal(title string, message interface{}) {
	GetOutput(p).Print(PLevelNormal, title, message)
}

// ToUnneccessary print as unneccesssary
func (p *PrintWrapper) ToUnneccessary(title string, message interface{}) {
	GetOutput(p).Print(PLevelUnneccessary, title, message)
}
