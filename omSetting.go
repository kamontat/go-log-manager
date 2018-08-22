package om

// FormatMessage is function to format message
type FormatMessage func(*Output, *Level, string, interface{}) string

// Setting is setting object for Output
type Setting struct {
	name   string
	Format FormatMessage
	Type   *OutputType
	Level  *Level
	To     *OutputTo
	Color  bool
}

// Name of the setting
func (s *Setting) Name() string {
	return s.name
}

// SetAsDefaultName will set current setting to by default value
func (s *Setting) SetAsDefaultName() {
	s.name = "default"
}

// SetAsDefaultType will set current setting to by default value
func (s *Setting) SetAsDefaultType() {
	s.Type = Logger
}

// SetMaximumLevel will update maximum level of this setting
func (s *Setting) SetMaximumLevel(level *Level) *Setting {
	s.Level = level
	return s
}

// SetAsDefaultLevel will set current setting to by default value
func (s *Setting) SetAsDefaultLevel() {
	if s.Type == Printer {
		s.Level = PLevelNormal
	} else {
		s.Level = LLevelInfo
	}
}

// SetOutputTo will update output of this setting
func (s *Setting) SetOutputTo(to *OutputTo) *Setting {
	s.To = to
	return s
}

// SetAsDefaultOutputTo will set current setting to by default value
func (s *Setting) SetAsDefaultOutputTo() {
	var name = "default"
	if s.name != "" {
		name = s.name
	}

	fileName := "/tmp/output/" + name

	if s.Type == Printer {
		fileName += ConstrantPrintSuffixName
	} else {
		fileName += ConstrantLogSuffixName
	}

	s.To = &OutputTo{Stdout: true, File: true, FileName: fileName}
}

// SetColor will update is color of this setting
func (s *Setting) SetColor(color bool) *Setting {
	s.Color = color
	return s
}

// SetAsDefaultColor will set current setting to by default value
func (s *Setting) SetAsDefaultColor() {
	s.Color = true
}

// SetAsDefaultFormat will set current setting to by default value
func (s *Setting) SetAsDefaultFormat() {
	if s.Type == Logger {
		s.Format = DefaultLoggerFormatMessage
	} else if s.Type == Printer {
		s.Format = DefaultPrinterFormatMessage
	}
}

// MarkDefault will mark all nil value to default value
func (s *Setting) MarkDefault() {
	if s.name == "" {
		s.SetAsDefaultName()
	}

	if s.Type == nil {
		s.SetAsDefaultType()
	}

	if s.Level == nil {
		s.SetAsDefaultLevel()
	}

	if s.To == nil {
		s.SetAsDefaultOutputTo()
	}

	if s.Format == nil {
		s.SetAsDefaultFormat()
	}
}

// DefaultSetting will create default setting.
// Name: default
// Color: true
// Type: Logger
// Level: Info
// To: Stdout, File (/tmp/output/default-time.log)
// Format: DefaultLoggerFormatMessage
func DefaultSetting() *Setting {
	s := &Setting{
		Color: true,
	}
	s.MarkDefault()
	return s
}
