package log

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

const (
	levelInfo  = "[INFO]"
	levelDebug = "[DEBUG]"
	levelWarn  = "[WARN]"
	levelError = "[ERROR]"
	levelFatal = "[FATAL]"
)

func Info(v ...interface{}) {
	plog(levelInfo, v...)
}

func Infof(format string, v ...interface{}) {
	plogf(levelInfo, format, v...)
}

func Debug(v ...interface{}) {
	plog(levelDebug, v...)
}

func Debugf(format string, v ...interface{}) {
	plogf(levelDebug, format, v...)
}

func Warn(v ...interface{}) {
	plog(levelWarn, v...)
}

func Warnf(format string, v ...interface{}) {
	plogf(levelWarn, format, v...)
}

func Error(v ...interface{}) {
	plog(levelError, v...)
}

func Errorf(format string, v ...interface{}) {
	plogf(levelError, format, v...)
}

func Fatal(v ...interface{}) {
	plog(levelFatal, v...)
	os.Exit(1)
}

func Fatalf(format string, v ...interface{}) {
	plogf(levelFatal, format, v...)
	os.Exit(1)
}

func plog(level string, v ...interface{}) {
	plogf(level, "", v...)
}

func plogf(level, format string, v ...interface{}) {
	if level == levelInfo {
		log.Print(level + " " + fmt.Sprintln(v...))
	} else {
		_, file, line, _ := runtime.Caller(3)
		loc := fmt.Sprintf("[%s:%d]", file, line)
		msg := fmt.Sprintln(v...)
		if len(msg) > 0 {
			msg = msg[0 : len(msg)-1]
		}
		log.Print(level + " " + msg + " " + loc)
	}
}
