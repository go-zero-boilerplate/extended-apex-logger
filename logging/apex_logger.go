package logging

import (
	"fmt"
	"strings"

	apex "github.com/francoishill/log"

	"github.com/go-zero-boilerplate/extended-apex-logger/utils/debug"
)

//NewApexLogger creates a new logger
func NewApexLogger(level apex.Level, handler apex.Handler, apexEntry *apex.Entry) Logger {
	//TODO: This set level and handler is global, this should not happen inside `NewApexLogger`
	apex.SetLevel(level)
	apex.SetHandler(handler)

	return &apexLogger{
		level:          level,
		handler:        handler,
		apexEntry:      apexEntry,
		errStackTraces: true, //TODO: Is this fine by default?
	}
}

type apexLogger struct {
	level          apex.Level
	handler        apex.Handler
	apexEntry      *apex.Entry
	errStackTraces bool
}

func (l *apexLogger) Emergency(s string) {
	l.apexEntry.Fatalf(s)
}
func (l *apexLogger) Alert(s string) {
	l.apexEntry.Errorf(s)
}
func (l *apexLogger) Critical(s string) {
	l.apexEntry.Errorf(s)
}
func (l *apexLogger) Error(s string) {
	l.apexEntry.Errorf(s)
}
func (l *apexLogger) Warn(s string) {
	l.apexEntry.Warnf(s)
}
func (l *apexLogger) Notice(s string) {
	l.apexEntry.Warnf(s)
}
func (l *apexLogger) Info(s string) {
	l.apexEntry.Infof(s)
}
func (l *apexLogger) Debug(s string) {
	l.apexEntry.Debugf(s)
}
func (l *apexLogger) Trace(s string) LogTracer {
	return l.apexEntry.Trace(fmt.Sprintf(s))
}
func (l *apexLogger) TraceDebug(s string) LogDebugTracer {
	return &localLogDebugTracer{l.apexEntry.TraceLevel(apex.DebugLevel, fmt.Sprintf(s))}
}

func (l *apexLogger) WithError(err error) Logger {
	newEntry := l.apexEntry.WithError(err)
	if l.errStackTraces {
		stack := strings.Replace(debug.GetFullStackTrace_Normal(false), "\n", "\\n", -1)
		newEntry = newEntry.WithField("stack", stack)
	}
	return NewApexLogger(l.level, l.handler, newEntry)
}

func (l *apexLogger) WithField(key string, value interface{}) Logger {
	return NewApexLogger(l.level, l.handler, l.apexEntry.WithField(key, value))
}

func (l *apexLogger) WithFields(fields map[string]interface{}) Logger {
	return NewApexLogger(l.level, l.handler, l.apexEntry.WithFields(apex.Fields(fields)))
}

func (l *apexLogger) DeferredRecoverStack(debugMessage string) {
	if r := recover(); r != nil {
		logger := l.WithField("recovery", r).WithField("debug", debugMessage)
		stack := strings.Replace(debug.GetFullStackTrace_Normal(false), "\n", "\\n", -1)
		logger = logger.WithField("stack", stack)
		logger.Alert("Unhandled panic recovered")
	}
}
