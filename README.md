# extended-apex-logger
An extended version (still using apex undelying) of the apex logger

## Getting started

main.go:

```
package main

import (
	"os"

	apex "github.com/francoishill/log"

	"github.com/go-zero-boilerplate/extended-apex-logger/logging"
	"github.com/go-zero-boilerplate/extended-apex-logger/logging/text_handler"
)

func getLogger() logging.Logger {
	level := apex.DebugLevel
	loggerFields := apex.Fields{}
	apexEntry := apex.WithFields(loggerFields)

	logHandler := text_handler.New(os.Stdout, os.Stderr, text_handler.DefaultTimeStampFormat, text_handler.DefaultMessageWidth)
	return logging.NewApexLogger(level, logHandler, apexEntry)
}

func main() {
	logger := getLogger()
	logger.Info("Hello world")
}
```