package main

import (
	//"log/syslog"
	//"log"
	//"fmt"
	"fmt"
)

func main() {
	//priority := syslog.LOG_LOCAL3 | syslog.LOG_NOTICE
	//flags := log.Ldate | log.Lshortfile
	//logger, err := syslog.NewLogger(priority, flags)

	/*logger, err := syslog.New(syslog.LOG_LOCAL3, "antony")
	if err != nil {
		panic("Cannot attache to syslog")
	}
	defer logger.Close()

	logger.Debug("Debug message.")
	logger.Warning("Warning message.")

	// Runtime and Runtime/Debug - packages gives functions to analyze memory and other stack details
	// runtime.Stack()
	// debug.PrintStack()

	*/

	fmt.Println("Test")
}
