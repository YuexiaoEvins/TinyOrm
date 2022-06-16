package log

import (
	"io/ioutil"
	"log"
	"os"
	"sync"
)

var (
	errorLog = log.New(os.Stdout, "\033[31m[error]\033[0m ", log.LstdFlags|log.Lshortfile)
	infoLog  = log.New(os.Stdout, "\033[34m[info ]\033[0m ", log.LstdFlags|log.Lshortfile)
	loggers  []*log.Logger
	mu       sync.Mutex
)

var (
	Error  = errorLog.Panicln
	Errorf = errorLog.Printf
	Info   = infoLog.Panicln
	Infof  = infoLog.Printf
)

const (
	InfoLevel = iota
	ErrorLevel
	Disables
)

func SetLogLevel(level int) {
	mu.Lock()
	defer mu.Unlock()

	for _, logger := range loggers {
		logger.SetOutput(os.Stdout)
	}

	if InfoLevel < level {
		infoLog.SetOutput(ioutil.Discard)
	}

	if ErrorLevel < level {
		errorLog.SetOutput(ioutil.Discard)
	}
}
