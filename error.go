package om

// import (
// 	"os"

// 	logger "github.com/kamontat/go-log-manager/logger"
// )

// // Exitable is object to exit program after error
// type Exitable struct {
// 	code int
// }

// func (e *Exitable) Exit() {
// 	os.Exit(e.code)
// }

// func (e *Exitable) ExitByCode(code int) {
// 	os.Exit(code)
// }

// type Exception struct {
// 	logger *logger.Logger
// }

// var exception *Exception

// func StartError() *Exception {
// 	if exception == nil {
// 		exception = &Exception{
// 			logger: GetLogger(),
// 		}
// 	}
// 	return exception
// }

// func (e *Exception) PrintError(title string, message string) {
// 	e.logger.Print(LLevelError, title, message)
// }

// func (e *Exception) ThrowOptionMissing(option string) Exitable {
// 	e.PrintError("Option miss", option+" is missing and it required.")
// 	return Exitable{code: 11}
// }

// func (e *Exception) ThrowArgumentMissing(argument string) Exitable {
// 	e.PrintError("Argument miss", argument+" is missing and it required.")
// 	return Exitable{code: 12}
// }

// func (e *Exception) ThrowDataMissing(data string, from string) Exitable {
// 	e.PrintError("Data miss", data+" is missing from "+from)
// 	return Exitable{code: 10}
// }

// func (e *Exception) ThrowGlobalConfigError(message string) Exitable {
// 	e.PrintError("global config error", message)
// 	return Exitable{code: 20}
// }
