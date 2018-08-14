package om

import (
	"fmt"
	"time"
)

var (
	// LogName is file of log
	LogName = fmt.Sprintf("%s-%d.log", "ndd", time.Now().Unix())
	// PrintName is file of printer
	PrintName = fmt.Sprintf("%s-%d.print", "ndd", time.Now().Unix())

	// DateFormat format of date in logger
	DateFormat = "02/Jan/06, 03:04:05.000"
)
